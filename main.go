package main

import (
	"encoding/json"
	"fmt"
	"moter"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/bamchoh/pasori"
	rpio "github.com/stianeikeland/go-rpio"
)

const (
	LockPin   = 23
	UnlockPin = 24
)

var (
	VID uint16 = 0x054C // SONY
	PID uint16 = 0x06C3 // RC-S380
)

type Config struct {
	UnlockAngle int `json:"unlockAngle"`
	LockAngle   int `json:"lockAngle"`
	SleepSec    int `json:"sleepSec"`
	AutoLockSec int `json:"autoLockSec"`
}

type Users map[string]string

func main() {
	config := loadConfig()
	users := loadUsers()

	idm, err := pasori.GetID(VID, PID)
	if err != nil {
		panic(err)
	}
	fmt.Println(idm)

	go button(config)

	for {
		idm := getIdm(nfc())
		unlockUser, ok := users[idm]
		if ok {
			unlock(config)
			fmt.Printf("***Welcome back %s!***\n", unlockUser)

			fmt.Printf("Wait %dsec...", config.AutoLockSec)
			time.Sleep(time.Duration(config.AutoLockSec) * time.Second)
			fmt.Println()

			lock(config)
		}
	}
}

func loadConfig() Config {
	file, _ := os.ReadFile("servoMoter.conf")
	var config Config
	json.Unmarshal(file, &config)
	return config
}

func loadUsers() Users {
	file, _ := os.ReadFile("user.conf")
	var users Users
	json.Unmarshal(file, &users)
	return users
}

func nfc() string {
	out, err := exec.Command("python", "/path/to/nfcpy/examples/tagtool.py").Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

func getIdm(text string) string {
	re := regexp.MustCompile(`ID=(.*?)\s`)
	match := re.FindStringSubmatch(text)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func unlock(config Config) {
	fmt.Println("Unlocking")
	moter.RunMoter(fmt.Sprint(config.UnlockAngle))
	time.Sleep(time.Duration(config.SleepSec) * time.Second)
	moter.RunMoter("0")
	fmt.Println("Unlocked")
}

func lock(config Config) {
	fmt.Println("Locking")
	moter.RunMoter(fmt.Sprint(config.LockAngle))
	time.Sleep(time.Duration(config.SleepSec) * time.Second)
	moter.RunMoter("0")
	fmt.Println("Locked")
}

func button(config Config) {
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	lockPin := rpio.Pin(LockPin)
	unlockPin := rpio.Pin(UnlockPin)

	lockPin.Input()
	lockPin.PullUp()

	unlockPin.Input()
	unlockPin.PullUp()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()

	for {
		if lockPin.Read() == rpio.Low {
			fmt.Println("lock button pushed.")
			lock(config)

		} else if unlockPin.Read() == rpio.Low {
			fmt.Println("unlock button pushed.")
			unlock(config)
			time.Sleep(time.Second)
			lock(config)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

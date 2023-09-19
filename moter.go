package moter

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const (
	servoPin = 18
)

func RunMoter(moveDeg int) {
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rpio.Close()

	pin := rpio.Pin(servoPin)
	pin.Mode(rpio.Pwm)
	pin.Freq(1920)
	pin.DutyCycle(0, 1024)

	fmt.Println(moveDeg)

	pin.DutyCycle(uint32(moveDeg), 1024)

	time.Sleep(1 * time.Second)
}

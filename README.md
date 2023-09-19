# zaregoto_door
Pasori S-330以降での動作を想定  

現行nfcpyだけ外部で呼んでいる，追々移植予定。  

## install
### install pasori-go
```shell
sudo apt install libusb-1.0-0 libusb-1.0-0-dev
```

### Start
```shell
cd z_door2.0
go run main.go
```

### GPIO構成
```shell
SERVO_PWM = gpio18
UNLOCK_BUTTON_PIN #gpio23
LOCK_BUTTON_PIN #gpio24
```
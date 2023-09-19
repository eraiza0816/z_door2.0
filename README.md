# zaregoto_door
Pasori S-330以降での動作を想定  

現行nfcpyだけ外部で呼んでいる，追々移植予定。  

## install

### install nfcpy
```shell
sudo apt install python-pip
sudo pip install -U nfcpy
sudo python -m nfc
```

```shell
** found usb:054c:06c3 at usb:006:002 but it's already used
-- scan sysfs entry at '/sys/bus/usb/devices/6-1:1.0/'
-- the device is used by the 'port100' kernel driver
-- this kernel driver belongs to the linux nfc subsystem
-- you can remove it to free the device for this session
   sudo modprobe -r port100
-- and blacklist the driver to prevent loading next time
   sudo sh -c 'echo blacklist port100 >> /etc/modprobe.d/blacklist-nfc.conf'
I'm not trying serial devices because you haven't told me
-- add the option '--search-tty' to have me looking
-- but beware that this may break other serial devs
```

言われたとおり  
` sudo modprobe -r port100 `
をしたあとに  
` sudo python -m nfc `
をしたらメッセージが変わっていた。  
` ** found SONY RC-S380/P NFC Port-100 v1.11 at usb:006:002 `

` sudo sh -c 'echo blacklist port100 >> /etc/modprobe.d/blacklist-nfc.conf' `

```shell
vi ~/nfcpy/nfcpy/nfc/llcp/sec.py
開いて191、192、201行をコメントアウト

# self.crypto.ECDH_OpenSSL.restype = c_void_p
# self.crypto.ECDH_set_method.restype = c_int
# self.crypto.EVP_CIPHER_CTX_init.restype = None
```

### Start
```shell
cd z_door2.0  
```

### GPIO構成
```shell
SERVO_PWM = gpio18
UNLOCK_BUTTON_PIN #gpio23
LOCK_BUTTON_PIN #gpio24
```
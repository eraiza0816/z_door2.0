# Zaregoto_door2.0
`nfcpy`は廃止されました。（時代を逆行）  
S330で動くことを期待します。  
`libpafe` `libpafe-ruby` を使います。
以下未修正のため読んでも意味ないです。

# install

## install ServoBlaster
```
git clone git://github.com/richardghirst/PiBits.git
cd PiBits/ServoBlaster/user
make
sudo make install
```

## install nfcpy
`sudo apt install python-pip`  
`pip install -U nfcpy`

`sudo python -m nfc`  
```
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

```
vim ~/nfcpy/nfcpy/nfc/llcp/sec.py
開いて191、192、201行をコメントアウト

# self.crypto.ECDH_OpenSSL.restype = c_void_p
# self.crypto.ECDH_set_method.restype = c_int
# self.crypto.EVP_CIPHER_CTX_init.restype = None
```
# Start
`cd ///zaregoto_door`  
`ruby controler.rb`

# Files
|file_names|usages|
|--|--|--|
|authed_cards.rb|Saved IDm info|
|controler.rb|controler|
|servo.rb|Saved servo motor info|

# Log
[hire](https://docs.google.com/spreadsheets/d/16dnM_fcZhj-seAavndLCHZvxERE-6B52Ia5F663mC1Q/edit#gid=0)

# Zaregoto_door2.0
`nfcpy`は廃止されました。（時代を逆行）  
S330で動くことを期待します。  
`libpafe` `libpafe-ruby` を使います。  
--以下修正中。--

# install

## install ruby
`sudo apt install ruby ruby-dev`

## install ServoBlaster
```
git clone git://github.com/richardghirst/PiBits.git
cd PiBits/ServoBlaster/user
make
sudo make install
```

## install libpafe
```
git clone https://github.com/rfujita/libpafe.git  
cd libpafe
./configure
make
make install
```  
うまく行かないときは **sudo**  つけると動いたりする。

## install libpafe
```
git clone https://github.com/rfujita/libpafe-ruby.git  
cd libpafe-ruby
ruby extconf.rb
make
make install
```

# Test
**sudo** 付けまくってるので `ruby`の実行時もsudoでやる。

```
cd libpafe-ruby
sudo ruby felica_dump.rb
```

# Start
`cd ///zaregoto_door`  
`ruby controler.rb`

# Files
|file_names|usages|
|--|--|--|
|controler.rb|controler|
|servo.rb|Saved servo motor info|

# Log
[hire](https://docs.google.com/spreadsheets/d/16dnM_fcZhj-seAavndLCHZvxERE-6B52Ia5F663mC1Q/edit#gid=0)

require './authed_cards.rb'
require './servo.rb'

def nfc()
    `sudo python ~/tmp/pasori/zaregoto_door/nfcpy/examples/tagtool.py`
end

def idm(text)
    m = text.match(/ID=(.*?)\s/)
    idm =m[1]
    print("IDm = #{idm}\n")
    return idm
end

def unlock()
    print("Unlocking\n")
    `echo #{SERVO_PIN}=#{UNLOCK_ANGLE}% > /dev/servoblaster`
end

def lock()
    print("Locking\n")
    `echo #{SERVO_PIN}=#{LOCK_ANGLE}% > /dev/servoblaster`
end

def log(unlock_user,operation)
  File.open("log.txt","a") do |f|
    f.print("#{Time.now},#{unlock_user},#{operation}\n")
  end
  print("\n#{Time.now},#{unlock_user},#{operation}\n")
end

loop do
    idm = idm(nfc)
    unlock_user = USERS.key(idm)

    unless unlock_user == nil
        print("***Welcome back #{unlock_user}!***\n")
        unlock

        print("Wait #{AUTO_LOCK_SEC}sec...")
        sleep(AUTO_LOCK_SEC)
        print("\n")

        lock
        puts("locked\n")
        %x[curl -X POST -H "Content-Type: application/json" -d '{"value1":"#{unlock_user}","value2":"control"}' https://maker.ifttt.com/trigger/keyLogAddSPsheet/with/key/bJh5tK010oplTRDxJ9zu2K]


    else
        print("***Illegal user***\n")
        %x[curl -X POST -H "Content-Type: application/json" -d '{"value1":"#{idm}","value2":"NO AUTHORITY"}' https://maker.ifttt.com/trigger/keyLogAddSPsheet/with/key/bJh5tK010oplTRDxJ9zu2K]


    end

    print("Please wait reader restart...\n")
end

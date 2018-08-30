#p("hekko")
#{}%x[sudo ruby ./felica_dump.rb]
#{}`sudo ruby ./felica_dump.rb`


puts(`vcgencmd measure_temp`)

def cpu_temp(text)
 cpu_temp = %x[vcgencmd measure_temp]
end
puts(cpu_temp)

# def cpu_temp()
#   cpu_temp = 0
# end
#
# if cpu_temp == text.match(vcgencmd:)

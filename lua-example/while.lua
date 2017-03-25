i = 0
while(i < 10)
do
	repeat
		i = i + 1
		if(i < 5) then
			break
		end
		print("break")
	until true
	print("until true")
end
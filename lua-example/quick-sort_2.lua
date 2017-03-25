
function printArray(arr1, s, e)
	for i= s, e do
	   print(arr1[i])
	end
end

-- i inclusive, j inclusive
function quick_sort(a, low, high)
	if(low >= high) then
		return
	end
	
    local i = low
	local pivot = a[high]
	local j = low - 1
	
	while(j < high)
	do
		j = j + 1
		if(a[j] < pivot) then
			t = a[i]
			a[i] = a[j]
			a[j] = t
			i = i + 1
		end
	end
	
	t = a[i]
	a[i] = pivot
	a[high] = t
		
	quick_sort(a, low, i - 1)
	quick_sort(a, i + 1, high)
end

arr = {12, 2, 23, 3, 24 ,34, 9 ,1, 32, 11}

quick_sort(arr, 1, 10)

print("final");
printArray(arr, 1, 10);



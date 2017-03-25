
function printArray(arr1, s, e)
	for i= s, e do
	   print(arr1[i])
	end
end

-- i inclusive, j inclusive
function quick_sort(a, i, j)
	if(i >= j) then
		return
	end
	abc = 3
	
    local originI = i
	local originJ = j
	local pivot = a[originI]
	while(i < j)
	do
		while(a[j] >= pivot and j > i)
		do
			j = j - 1
		end
		a[i] = a[j]
		while(a[i] <= pivot and i < j)
		do
			i = i + 1
		end
		a[j] = a[i]
	end
	
	a[i] = pivot
	
	quick_sort(a, originI, i - 1)
	quick_sort(a, i + 1, originJ)
end

arr = {12, 2, 23, 3, 24 ,34, 9 ,1, 32, 11}

quick_sort(arr, 1, 10)

print("final");
printArray(arr, 1, 10);


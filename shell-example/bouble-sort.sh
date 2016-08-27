#!/bin/bash
array=$@
IFS=' '
array=($array)
echo origin array: $@
length=${#array[@]}
echo length is $length
for (( i=0 ; i<$length; i=i+1 ))
do
	for (( j=0; j<$((length-i-1)); j=j+1 ))
	do
		if (( ${array[$j]} > ${array[$((j+1))]} )); then
			temp=${array[$j]}
			array[$j]=${array[$((j+1))]}
			array[$j+1]=$temp
		fi

	done
done

for (( i=0 ; i<$length; i=i+1))
do
	echo ${array[$i]}
done


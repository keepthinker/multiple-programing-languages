#!/bin/sh

#num=20

echo '\nenter a number between 0 and 20:'

read num

factorial(){
	a=0
	result=1
	while [ "$a" -lt "$1" ]
	do
		a=`expr $a + 1`
		result=`expr $result \* $a`
		echo $result
	done
	echo "$1! equals $result"
}

if [ $num -gt 20 ]
then
	echo 'num must be in [0,20]'
else
	factorial $num
fi

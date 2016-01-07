#!/bin/sh

rawStr = raw_input("input(separated by ',') : ");

strArr = rawStr.split(",");

def sort_string_array(strArr):
	"sort array desc alphabetically"
	length = len(strArr);
	for i in range(0, length - 1, 1):
		for j in range(0, length - i - 1):
			if strArr[j] > strArr[j + 1]:
				temp =  strArr[j + 1];
				strArr[j + 1] = strArr[j];
				strArr[j] = temp;
				print strArr,i, j, strArr[j], strArr[j+1];


sort_string_array(strArr);
print(strArr)


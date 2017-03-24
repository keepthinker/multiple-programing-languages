
expectedYearNum = input("year: ")
year = 10
if(len(expectedYearNum) != 0):
  year = int(expectedYearNum)

base = 100000

general_incr_ratio = 1.2

bull_incr_ratio = 3

i = 0

income = 0

bull_cycle = 8

acutual_incr_ratio = general_incr_ratio

while(i < year):
  if(i % bull_cycle == 6):
    actual_incr_ratio = bull_incr_ratio
  else:
    actual_incr_ratio = general_incr_ratio
  income =  (income + base) * actual_incr_ratio
  print("year--" + str(i) + ":" + str(income))
  i += 1

print(income)
      

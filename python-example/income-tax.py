import sys

income = int(sys.argv[1])

base = 3500
level1500 = 1500 * 0.03
level4500 = level1500 + 3000 * 0.1
level9000 = level4500 + 4500 * 0.2
level35000 = level9000 + 26000 * 0.25
level55000 = level35000 + 20000 * 0.3
level80000 = level55000 + 25000 * 0.35

tax = 0
rest = income - base
if(rest > 0 and rest <= 1500):
  tax = rest * 0.3
elif(rest > 1500 and rest <= 4500):
  rest -= 1500
  tax += level1500 + rest * 0.10
elif(rest > 4500 and rest <= 9000):
  rest -= 4500
  tax += level4500 + rest * 0.20;
elif(rest > 9000 and rest <= 35000):
  rest -= 9000
  tax += level9000 + rest * 0.25
elif(rest > 35000 and rest <= 55000):
  rest -= 35000
  tax += level35000 + rest * 0.30
elif(rest > 55000 and rest <= 80000):
  rest -= 55000
  tax += level55000 + rest * 0.35
elif(rest > 80000):
  rest -= 80000
  tax += level80000 + rest * 0.45
else:
  tax = 0;

oldTax = tax

print "before 2018, tax is ", tax

tax = 0
yearIncome = income * 12
base2019 = 5000
rest = yearIncome - base2019 * 12
#print "rest:", rest
if(rest > 0 and rest <= 36000):
  tax = rest * 0.03
elif(rest > 36000 and rest <= 144000):
  tax = rest * 0.1 - 2520
elif(rest > 144000 and rest <= 300000):
  tax = rest * 0.2 - 16920
elif(rest > 300000 and rest <= 420000):
  tax = rest * 0.25 - 31920
elif(rest > 420000 and rest <= 660000):
  tax = rest * 0.3 -52920
elif(rest > 660000 and rest <= 960000):
  tax = rest * 0.35 - 85920
elif(rest > 960000):
  tax = rest * 0.45 - 181920
else:
  tax = 0

tax = tax / 12
print "after 2018, tax is ", tax

print "after 2018, tax is reduced by ", oldTax - tax

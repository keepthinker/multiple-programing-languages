def func(param, comment = 'my comment'):
    print param, comment
func('asd', 'my')
func(param = "asd", comment = 'comment')

def print_strs(first, *params):
    "print str concatenated"
    for elem in params:
        first += elem
    print first
    
print_strs('asdf', 'gh', 'ijk')

sum = lambda a1, a2: a1 + a2
print sum(1,2)

print_strings = print_strs;
print_strings('a', 'b')

list = [1,2,3,4]
i = 0
len1 = len(list)
while(i < len1):
    print list[i]
    i += 1

tuple = ('a', 'b', 'c', 'd');
i = 0
len1 = len(tuple)
print "length of tuple : ", len1
while(i < len1):
    print tuple[i]   
    i += 1

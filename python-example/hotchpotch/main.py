pyg = 'ay'
word = raw_input("enter a english word : ")
i = 0;
word_len = len(word)
isAllowd = True;
while i < word_len:
    c = word[i];
    if not word.isalpha():
        isAllowd = False;
        print "not a english word %s" % c
        break;
    i += 1
    
if( isAllowd == True):
    first_letter = word[0];
    word = word[1:len(word)]
    word += first_letter + pyg
    print word;

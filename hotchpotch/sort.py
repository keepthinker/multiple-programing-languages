class Sort:
    'For sorting array'
    def __init__(self):
        print 'initialized'
        return

    'Bubble sort'
    @staticmethod
    def sort(array):
        length = len(array)
        i = length;
        while(i > 0):
            j = 1
            while(j < i):
                if(array[j - 1] > array[j]):
                    temp = array[j];
                    array[j] = array[j - 1]
                    array[j - 1] = temp
                j += 1
            i -= 1
       

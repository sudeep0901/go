import random

def weekdays():
    days = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']
    # i=days.index(day) # get the index of the selected day
    # d1=days[i:] #get the list from an including this index
    # d1.extend(days[:i]) # append the list form the beginning to this index
    return days[0:5]

timetable = {}

def init_timetable(dow, sub):
    timetable[dow] = sub
    print(dow, sub)

dow_arr = weekdays()
classes = list(range(0,10))
rooms = list(range(0,35))
subjects = list(range(0,7))
teachers = list(range(0,35))
# subjects = ['s' + str(s) for s in subjects]
teachers = ['t' + str(s) for s in teachers]

print(dow_arr)
print(subjects)
print(teachers)


for sub in subjects:
    print(sub)
    for index, dow in enumerate(dow_arr):
        print(dow)
        rand_seq_subjects =  random.sample(subjects, len(subjects))
        print(rand_seq_subjects)
        if(index == 0):
            init_timetable(dow, rand_seq_subjects)
        else:
            curr_pos = len(rand_seq_subjects) - 1
            temp_element =  rand_seq_subjects[curr_pos]

            for elem in rand_seq_subjects:
                init_timetable(dow, rand_seq_subjects)

print(timetable.keys())
print(timetable.values())

s = 'babd'

if len(s) == 1:
    print("palindrome string")
    exit(0)
long_palindrome = ''
for i in range(len(s)):
    for j in range(i + 1,len(s)):
        if j > 2 and (j - i) %2  != 0:
            continue;
        # print("i:", i, "j:",j)
        isPalin = True
        ts = s[i:j + 1]
        for v in range(len(ts)):
            print(ts)
            print(v,ts[v], ts[-v-1])
            if ts[v] != ts[-v -1]:
                isPalin = False
                break
        if (isPalin) and len(long_palindrome) < len(ts):        
            long_palindrome = ts

print(palin)
print(long_palindrome)
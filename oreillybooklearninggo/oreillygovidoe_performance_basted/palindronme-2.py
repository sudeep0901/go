s = "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffgggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"

print(len(s))
print(s == s[::-1])
if len(s) == 1:
    print("palindrome string")
    exit(0)
long_palindrome = ''
for i in range(len(s)):
    print(i)
    for j in range(i,len(s)):
        # if j > 2 and (j - i) %2  != 0:
            # continue;
        # print("i:", i, "j:",j)
        isPalin = True
        ts = s[i:j + 1]
        for v in range(len(ts)):
            # print(ts)
            # print(v,ts[v], ts[-v-1])
            if ts[v] != ts[-v -1]:
                isPalin = False
                break
        if (isPalin) and len(long_palindrome) < len(ts):        
            long_palindrome = ts

print(long_palindrome)
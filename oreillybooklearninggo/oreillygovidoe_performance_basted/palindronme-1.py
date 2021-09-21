s = "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa"

if len(s) == 1:
    print("palindrome string")
    exit(0)
long_palindrome = ''
for i in range(len(s)):
    for j in range(i + 1,len(s)):
        # if j > 2 and (j - i) %2  != 0:
            # continue;
        isPalin = False
        ts = s[i:j + 1]
        for v in range(len(ts)):
            # print(ts)
            if ts == ts[::-1]:
                isPalin = True
        if (isPalin) and len(long_palindrome) < len(ts):        
            long_palindrome = ts

print(long_palindrome)
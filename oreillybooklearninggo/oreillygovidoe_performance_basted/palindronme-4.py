
def findLongsetPalidrome(s: str):
    if s == "":
        return s

    rev = s[::-1]

    dp = [[0 for i in range(len(s))] for j in range(len(s))]
    print(dp)
    max_len = 0
    max_end = 0

    for i in range(len(s)):
        for j in range(len(s)):
            print("s:", s[i], "i:", i, "rev:", rev[j], "j:", j)
            if s[i] == rev[j]:
                if i == 0 and j == 0:
                    print("Updating dp:")
                    dp[i][j] = 1
                else:
                    dp[i][j] = dp[i - 1][j-1] + 1
                print(dp[i][j])
                if dp[i][j] > max_len:
                    if i-dp[i][j]+1 == len(s)-1-j:
                        max_len = dp[i][j]
                        max_end = i
    print(dp, max_end, max_len)

    return s[max_end - max_len + 1: max_end + 1]


s = "babd"

print(findLongsetPalidrome(s))

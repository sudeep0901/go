
def findLongsetPalidrome(s: str):
    if s == "":
        return s
    print("Input string length:", len(s))
    dp = [[0 for i in range(len(s))] for j in range(len(s))]
    max_len = 0
    max_end = 0
    x, y = 0, 0
    for i in range(len(s) - 1, -1, -1):
        for j in range(i, len(s)):
            print("i:", i, "j:", j)

            if i == j:
                dp[i][j] = True
            elif s[i] == s[j]:
                print(s[i], s[j])
                if j - i == 1:
                    dp[i][j] = True
                else:
                    dp[i][j] = dp[i + 1][j - 1]
            if (dp[i][j] and j - i >= max_len):
                max_len = j - i
                x = i
                y = j

    print(dp, x, y)
    print(s[x: y - x + 1])
    return ''


s = "babd"

print(findLongsetPalidrome(s))

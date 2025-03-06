def length_of_longest_substring(s):
    char_set = set()
    left, max_len = 0, 0

    for right in range(len(s)):
        while s[right] in char_set:
            char_set.remove(s[left])
            left += 1
        char_set.add(s[right])
        max_len = max(max_len, right - left + 1)

    return max_len


s = "abcabcabc"
print(length_of_longest_substring(s))  # 输出 3

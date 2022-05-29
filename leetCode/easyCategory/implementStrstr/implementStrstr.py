# This is easy to implement in python since it has a built in function
class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        try:
            firstInstance = haystack.index(needle)
            return firstInstance
        except:
            return -1

haystackVar = "hello does value of needle exiiiist in this haystack ?"
needleVar = "ii"
strModified = Solution()
result = strModified.strStr(haystack=haystackVar, needle=needleVar)
if result == -1:
    print("'{0}' is not found in '{1}'".format(needleVar, haystackVar))
else:
    print("'{0}' is found in '{1}' at index: '{2}'".format(needleVar, haystackVar, result))
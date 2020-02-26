# Given nums = [2, 7, 11, 15], target = 9,
# Because nums[0] + nums[1] = 2 + 7 = 9,
# return [0, 1].

array = [2, 7, 11, 15]
target = 9

# What is a hash map?
# Data structure that maps a key => value.
# Benefits = O(1) because it does a hash of the key and uses that hash to look
# up the element.
#  my_dict={'Dave' : '001' , 'Ava': '002' , 'Joe': '003'}

# Dave => 09z345
# Ava => 82xz

# [09z, 82xz]

# O^n, constant time

# 1. Range over each element of array.
# 2. Subtract target - i (element of array)
# 3. insert result into hashmap...


def twoNumSum(array, target):
    # declare a dict/hashmap
    nums = {}
    for num in array:
        # 2, 7
        potentialMatch = target - num
        # potentialMatch = 9 - 2 = 7, 9 - 7 = 2
        if potentialMatch in nums:
            return [potentialMatch, num]
        else:
            nums[num] = True
            # nums{2 = True,}
    return []


print(twoNumSum(array, target))

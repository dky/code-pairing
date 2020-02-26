# Given nums = [2, 7, 11, 15], target = 9,
# Because nums[0] + nums[1] = 2 + 7 = 9,
# return [0, 1].

array = [2, 7, 11, 15]
target = 9


#  O^2 quadratic
def sumIndex(array, target):
    for i in range(len(array)):
        for j in range(i + 1, len(array)):
            if (array[i] + array[j] == target):
                print('index 1: ', 'index 2:', i, j)


sumIndex(array, target)

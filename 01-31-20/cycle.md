# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
    # node.next has a step of 1 , everytime you call node.next you move to the next node 
    # node.next.next has a step of 2, everything you call node.next.next is like calling node = node.next and then calling node.node.ext
    # [10] -> [20] -> [22] -> [40] head = [10]
    # head.next = [20]
    # head.next.next = [22]
    # head.next.next.next = [40]
    # head.next.next.next.next = None
    # node = head.next // node = [20]
    # node2 = node.next.next // node2 [40]
    # [10] -> [20] -> [22] -> [40]
    #          |_______________|
    # head.next.next.next.next = [20]
    # while node.next != null : display node ---> [10] [20] [22] [40] [20] [22] [40] ...... your cpu crashes
    # while node.next.next != null : { display node ; node = node.next.next } -> [10] [22] [20] [40] [22] [20] [40] ...... your cpu crashe
    # both speeds moving in one loop ( double pointer ) 
    # nodes: 			  [1] -> [2]  ->  [3] ->  [4]  ->  [5] 
    # trap to avoid : [10] -> [10] -> [10] -> [10] -> [10] 

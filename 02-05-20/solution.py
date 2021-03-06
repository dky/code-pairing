# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        slow = head
        fast = head

        while fast is not None and fast.next is not None:
            fast = fast.next.next
            slow = slow.next

        slow = self.reverse(slow)
        fast = head

        while slow is not None:
            if slow.val != fast.val:
                return False
            else:
                slow = slow.next
                fast = fast.next

        return True
    def reverseList(self, head: ListNode) -> ListNode:
        prev = None
        current = head 
        while current is not None: 
            next = current.next
            current.next = prev 
            prev = current 
            current = next
        head = prev
        
        return head

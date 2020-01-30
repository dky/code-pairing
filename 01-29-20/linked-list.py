# Given the following implmenent into linked list:

# HEAD => 1 => 2 => 3 => 4 => Null

# Singly linked list.

# Class node


class node:
    # Constructor
    def __init__(self, value=None, next=None):
        self.value = value
        self.next = next
        # print(value)

    # Read a linked list value
    # Never access value directly call func (accessor)
    def getValue(self):
        return self.value

    def next(self):
        return self.next

    def hasNext(self):
        # return 1 == 1
        return self.next != None


my_node_2 = node(2)
my_node = node(1, my_node_2)

while my_node.hasNext():
    print(my_node.getValue())
    my_node = my_node.next()

# print(my_node.getValue())
# print(my_node.next.getValue())
# print(my_node.hasNext())

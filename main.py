''' 
13/05/2022
Given the root to a binary tree, implement: 
- serialize(root), which serializes the tree into a string;
- deserialize(string), which deserializes the string back into the tree. '''

class Node :
  def __init__(self,value,left=None,right=None):
    self.value = value
    self.left = left
    self.right = right
    
# Serialize the last node as a string
def serializer(rootNode):
  print("> Loading",rootNode.value)
  if isinstance(rootNode.left, Node):
    print("+ Found left side node")
    return serializer(rootNode.left)
  elif rootNode.left == None:
    print("= This is final node")
    return rootNode.value
  return output

def deserializer (string :str):
  pass
    
node = Node('root', Node('left', Node('left.left',Node("left.left.left"))), Node('right'))
print(serializer(node))
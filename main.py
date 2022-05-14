''' 
Given the root to a binary tree, implement: 
- serialize(root), which serializes the tree into a string;
- deserialize(string), which deserializes the string back into the tree. '''

class Node :
  def __init__(self,value,left=None,right=None):
    self.value = value
    self.left = left
    self.right = right


def serializer (node):

  out = ""
  
  if type(node.left) == Node:
    out += serializer(node.left)
  else:
    out += node.left

  return out



  
    
node = Node('root', Node('left', Node('left.left')), Node('right'))
print( serializer (node) )
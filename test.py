word = "kitten"
reach = "sitting"


''' Calculate the shortest between the two: we will use that to iterate on the
longest'''

longest = reach if len(reach)>len(word) else word
shortest = reach if len(reach)<len(word) else word


'''Transform the longest into a list, so we can remove identical chars from
the two'''
longestToList = list(longest)

for c in shortest:
    if c in longestToList:
        longestToList.remove(c)


print(longestToList)

'''Now the list will contains the different letters from the two and the
missing ones. We just print the lenght of it'''
print(len(longestToList))




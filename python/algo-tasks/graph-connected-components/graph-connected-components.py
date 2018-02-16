"""
For a given graph, print the number of connected components
"""

graph = {}
connComps = {}
connCompsCnt = 0

def connect(node: str, identifier: str):
    if node in connComps:
        return

    connComps[node] = identifier
    for neighbour in graph[node]:
        connect(neighbour, identifier)

def readGraph():
    print("Input the count of nodes you are going to input")
    nodesCnt = int(input())
    print("Type the count of edges you are going to input")
    edgesCnt = int(input())

    print("input edges in the format {node} {neighbour}")
    for i in range(nodesCnt):
        graph[str(i)] = []

    for i in range(edgesCnt):
        nodeInfo = input().split(' ')
        node = nodeInfo[0]
        neighbour = nodeInfo[1]

        assert(node in graph)
        assert(neighbour in graph)
        assert(neighbour not in graph[node])
        assert(node not in graph[neighbour])

        graph[node].append(neighbour)
        graph[neighbour].append(node)

readGraph()
for node in graph:
    connect(node, node)

# 1. get values of dict
# 2. cast to set (keeps only distinct values)
# 3. get the length of the set
print(len(set(connComps.values())))

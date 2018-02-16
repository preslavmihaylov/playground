import re

graph = {}
topologicalSort = []
predecessorsCnt = {}

def readGraph():
    print("Input count of nodes")
    nodesCnt = int(input())

    print("Input node info in format {node}->{node1}, {node2}...")
    for i in range(nodesCnt):
        nodeInfo = input().split('->')
        node = nodeInfo[0]
        print(node)

        if not len(nodeInfo) == 1:
            neighbours = re.findall(r"[\w]+", nodeInfo[1])
        else:
            neighbours = []

        if node not in predecessorsCnt:
            predecessorsCnt[node] = 0

        graph[node] = neighbours
        for neighbour in neighbours:
            if neighbour not in graph:
                graph[neighbour] = []
            if neighbour not in predecessorsCnt:
                predecessorsCnt[neighbour] = 0

            predecessorsCnt[neighbour] += 1

def extractTopologicalSort():
    while True:
        # find minimum key based on the count of predecessors
        first = min(graph.keys(), key=lambda n:predecessorsCnt[n], default=None)
        if first is None:
            return

        topologicalSort.append(first)
        for neighbour in graph[first]:
            predecessorsCnt[neighbour] -= 1

        graph.pop(first)

readGraph()
extractTopologicalSort()

print("Topological sort:", ", ".join(topologicalSort))

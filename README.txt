LEM-IN

1. PROBLEM
2. SOLUTION
3. IMPLEMENTATION
4. CURIOSITIES
5. BIBLIOGRAPHY

1. PROBLEM

Suppose we're given a number of ants and a network of rooms connected by tunnels. One room is labelled "start" and another "end". Find a way to move all the ants from "start" to "end" in the smallest number of turns, subject to the following contraints: one ant per tunnel per turn, and one ant per room at the end of a turn except for "start" and "end" which can contain any number of ants. (See the 01-Edu Public Repo [0].)

2. SOLUTION

We first find a maximum flow (a largest set of compatible paths) through the corresponding undirected flow network with unit capacity on all nodes and edges.

There are several ways to do this. One is the Ford-Fulkerson method. As originally described for directed graphs, this works by defining an auxiliary network called the residual graph, having the same nodes and edges, except that the weight of each edge, known as the residula capacity on that edge, is set equal to capacity minus flow. Initially the flow is set to zero. Then, for as many iterations as possible, we find a provisional path with no cycles (loops) from "start" to "end" in the residual graph. When such a path is found, we send as much flow as possible along it (always 1 in our case) and adjust the residual capacities accordingly. In this way, we get to reverse flow where needed (cancel it out) if we make a less than optimal choice of path. Every new path increases the flow. When no more paths are possible, the flow is maximal.

[S] outlines Ford-Fulkerson and shows a way to adapt this technique to an undirected graph. This is relevant because we've assumed that ants can pass either way along the tunnels. Our instructions make no mention of a preferred direction. At first, we wondered if the order that the rooms are named in the list of connections might indicate a direction, but there are counterexamples among the audit solutions: In example02, a connection is listed as "3-2", but, in turn three, L2 moves from "2" to "3":

L1-3 L2-1
L2-2 L3-3 L4-1
L2-3 L4-2 L5-3 L6-1

In example05, a connection is listed as "D2-E2", but, in turn four, L4 moves from "E2" to "D2":

L1-A0 L4-B0 L6-C0
L1-A1 L2-A0 L4-B1 L5-B0 L6-C1
L1-A2 L2-A1 L3-A0 L4-E2 L5-B1 L6-C2 L9-B0
L1-end L2-A2 L3-A1 L4-D2 L5-E2 L6-C3 L7-A0 L9-B1

Ford-Fulkerson doesn't specify how the paths are to be found. If paths are found randomly, it will still work, but there are more specific algorithms with better-than-random choice of paths. We use one of these: Edmonds-Karp [W]. Edmonds-Karp is like Ford-Fulkerson except that instead of trying paths at random, it finds (at each step) a shortest valid path using breadth first search (BFS).

Note that Edmonds-Karp (and Ford-Fulkerson in general) doesn't place any capacity constraint on nodes. So we need to add a condition to prevent a new path in the residual graph from sharing a node with one of the existing paths of flow unless it also reverses the flow along an edge conncted to that node.

Maximum flow is optimal in the long run, as the number of ants increases without bound. By favouring shorter paths, Edmonds-Karp ensures that our solution also results in the smallest number of turns. The only exception might be if the number of ants was smaller than the value of the maximum flow and there were paths shorter than those in the particular maximum flow found. We don't want to lose short paths to gain more paths if there are enough short paths! (This would happen in example01, say, if there were less than three ants.) To eliminate this possibility, the program stops searching for paths after any iteration where the number of paths becomes to equal the number of ants.

3. IMPLEMENTATION

Aside from some error checking, the task is essentially divided into five functions:

* ParseNest parses the nest into structs of type Nest and Room.
* MaxFlow uses BFS to find paths according to Edmonds-Karp.
* PathCollector gathers these paths into a slice of items of struct type Path.
* SendAnts assigns ants to paths according to the scheme described by [D].
* PrintTurns formats the result in the style of the audit solutions.

Most important conceptually is MaxFlow. This function implements the Edmonds-Karp algorithm (i.e. Ford-Fulkerson with BFS), adapted to undirected graphs (per [S]) and streamlined to our case of unit capacity on all edges, but with the additional constraint of node capacity and the extra rule to stop when there's a path for every ant.

We implement the queue as a slice of (pointers to) rooms. The BFS fans out from "start" till a shortest route to "end" is found, subject to the residual capacity constraints. As the search moves on from node "u" to node "v", say, we set the "v.Predecessor" field equal to "u" to mark where we came from. The Predecessor field thus serves to mark which nodes have been visited during a particular iteration of the search for paths. Predecessor also signals when the "end" has been found because then "end.Predecessor != nil". This results in a linked list of rooms, which can now be traced back from "end" to "start" and "u.Flow[v]" set to "true" everywhere along the list, except where an edge previously had flow from "v" to "u" (i.e. v.Flow[u] was equal to "true"). In that case, the flow is cancelled out: both u.Flow[v] and v.Flow[u] are set to "false". It's these Flow fields that will remember the provisional paths after each step of the path search, while the Predecessor fields of all rooms are reset to "nil" at the start of the next iteration.

Future iterations of the path search revise and augment the flow, as described above. When no more paths can be found without breaking the capacity constraints or causing the number of paths to exceed the number of ants, PathCollector turns the resulting linked lists of flow into objects of struct type Path. The rooms belonging to each path, p, are stored in a slice in the p.Room field. The paths themselves are collected into a slice and ordered by length for ease of assigning the ants.

4. CURIOSITIES

Depending on the network and number of ants, there may exist optimal solutions with fewer-than-maximal paths. The audit answer for example05 is such a case. The nest is big enough and the number of ants small enough to achieve the smallest number of turns with only three paths. However, as the number of ants is increased, eventually these three tunnels require more turns than our maximal solution of four paths. Thus, with nine ants, both solutions take eight turns, but, with 99 ants, ours takes 30 turns, while theirs takes 38.

On the other had, although our program gives a solution with the smallest number of turns, it can happen that other, shorter paths are available for the first few ants, permitting a solution with just as few turns, but even fewer individual ant-moves. This is the case in example01, where the first ant to go to "h" can take one of the shorter paths, start-h-n-e-end or start-n-m-end, without blocking ants coming via "0" or "t", provided all other ants follow the three longer paths of the maximum flow.

5. BIBLIOGRAPHY

[0] 01-Edu: Public Repo [ https://github.com/01-edu/public/tree/master/subjects/lem-in ]. Accessed Jan. 1, 2023.

[D] Dawson J: Lem-in: Finding all the paths and deciding which are worth it [ https://medium.com/@jamierobertdawson/lem-in-finding-all-the-paths-and-deciding-which-are-worth-it-2503dffb893 ]. Nov. 19, 2019. Accessed Jan. 1, 2023.

[S] Schroeder J, Guedes AP, Duarte EP: Computing the Minimum Cut and Maximum Flow of Undirected Graphs
[ https://www.inf.ufpr.br/pos/techreport/RT_DINF003_2004.pdf ]. RT-DINF 003/2004. Accessed Jan. 1. 2023.

[W] Wikipedia: Edmonds-Karp algorithm [ https://en.wikipedia.org/wiki/Edmonds%E2%80%93Karp_algorithm ]. Apr. 14, 2022. Accessed Jan. 1. 2023.
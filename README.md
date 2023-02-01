# LEM-IN

1. [PROBLEM](#1-problem)
2. [SOLUTION](#2-solution)
3. [IMPLEMENTATION](#3-implementation)
4. [FURTHER NOTES](#4-further-notes)
5. [CURIOSITIES](#5-curiosities)
6. [BIBLIOGRAPHY](#6-bibliography)

## 1. PROBLEM

Suppose we're given a number of ants and a network of rooms connected by tunnels. One room is labelled `start` and another `end`. Initially all ants are in `start`. Find a way to move all the ants to `end` in the smallest number of turns, subject to the following contraints: one ant per tunnel per turn, and one ant per room at the end of a turn except for `start` and `end`, which can contain any number of ants. (See the 01-Edu Public Repo [^0].)

## 2. SOLUTION

We first start looking for a maximum flow (a largest set of compatible paths) through the corresponding undirected flow network with unit capacity on all nodes and edges.

There are several ways to do this. One is the Ford-Fulkerson method. As originally described for directed graphs, this works by defining an auxiliary network called the residual graph, having the same nodes and edges as the original graph, except that the weight of each edge, known as the residual capacity, is set equal to capacity minus flow. Initially the flow is set to zero. Then, for as many iterations as possible, we find a provisional path (known as an augmenting path) from `start` to `end` (with no cycles) in the residual graph. When such a path is found, we send as much flow as possible along it (always 1 in our case) and adjust the residual capacities accordingly. In this way, we get to push back flow where needed (cancel it out) if we make a less than optimal choice of path. Every new path increases the total flow, hence the name augmenting path. When no more paths are possible, the flow is maximal.

Schroeder, Guedes, and Duarte[^S] outline Ford-Fulkerson and show a way to adapt this technique to undirected graphs. This is relevant because we've assumed that ants can pass either way along the tunnels. Our instructions make no mention of a preferred direction. At first, we wondered if the order that the rooms are named in the list of connections might indicate a direction, but there are counterexamples among the audit solutions: In [example02](nests/audit_examples/example02), a connection is listed as `3-2`, but, in turn three, `L2` moves from `2` to `3`:

```
L1-3 L2-1
L2-2 L3-3 L4-1
L2-3 L4-2 L5-3 L6-1
```

In [example05](nests/audit_examples/example05), a connection is listed as `D2-E2`, but, in turn four, `L4` moves from `E2` to `D2`:

```
L1-A0 L4-B0 L6-C0
L1-A1 L2-A0 L4-B1 L5-B0 L6-C1
L1-A2 L2-A1 L3-A0 L4-E2 L5-B1 L6-C2 L9-B0
L1-end L2-A2 L3-A1 L4-D2 L5-E2 L6-C3 L7-A0 L9-B1
```

Ford-Fulkerson doesn't specify how the paths are to be found. If paths are found randomly, it will still work, but there are more detailed algorithms that follow the Ford-Fulkerson method except with better-than-random choice of paths. We use one of these: Edmonds-Karp [^W]. At each step, Edmonds-Karp finds a shortest valid path using breadth first search (BFS).

Note that Edmonds-Karp (and Ford-Fulkerson in general) doesn't place any capacity constraint on nodes. So we need to add a condition to prevent a new path in the residual graph from sharing a node with one of the existing paths of flow unless it also reverses the flow along an edge conncted to that node.

(Equivalently we could have substituted every node `u` with a pair of nodes, an entrance node and an exit node, connected them with a directed edge from entrance to exit, and replaced every edge connected to `u` with an incoming edge connected to the entrance node and an outgoing edge connected to the exit node.)

By favouring maximum flows with shorter paths, Edmonds-Karp finds a solution with the smallest number of turns PROVIDED THERE ARE ENOUGH ANTS. In some graphs, however, a maximum flow might include multiple longer paths that block a shorter path. In that event, below a certain number of ants, fewer but shorter paths are best. (See [nests/sneaky_examples/few.txt](nests/sneaky_examples/few.txt).) To eliminate this possibility, our program stops searching if more paths would actually increase the number of turns needed for the given amount of ants.

## 3. IMPLEMENTATION

Aside from some error checking, the task is essentially divided into five functions:

* [ParseNest](lem/parse_nest.go) parses the nest into structs of type [`Nest`](lem/structs.go) and [`Room`](lem/structs.go).
* [PathFinder](lem/path_finder.go) uses BFS to find paths according to Edmonds-Karp.
* It calls [PathCollector](lem/path_collector.go) to gather these paths into a slice of items of struct type [`Path`](lem/structs.go).
* Then it calls [SendAnts](lem/send_ants.go) to assign ants to paths according to the scheme described by Jamie Dawson[^D].
* Finally, [PrintTurns](lem/print_turns.go) formats the result in the style of the audit solutions.

Most important conceptually is `PathFinder`. This function implements the Edmonds-Karp algorithm (i.e. Ford-Fulkerson with BFS), adapted to undirected graphs (per Schroeder, Guedes, Duarte[^S]) and streamlined to our case of unit capacity on all edges, but with the additional constraint of node capacity and the extra rule to stop searching if more paths would increase the number of turns.

We implement the queue as a slice of (pointers to) rooms. The BFS fans out from `start` till a shortest route to `end` is found, subject to the residual capacity constraints. As the search moves on from node `u` to node `v`, say, we set the `v.Predecessor` field equal to `u` to mark where we came from. The Predecessor field thus serves to mark which nodes have been visited during a particular iteration of the search for paths. Predecessor also signals when the `end` has been found because then `end.Predecessor != nil`. This results in a linked list of rooms, which can now be traced back from `end` to `start` and `u.Flow[v]` set to `true` everywhere along the list, except where an edge previously had flow from `v` to `u` (i.e. `v.Flow[u]` was equal to `true`). In that case, the flow is cancelled out: both `u.Flow[v]` and `v.Flow[u]` are set to `false`. It's these Flow fields that will remember the provisional paths after each step of the path search, while the `Predecessor` fields of all rooms are reset to `nil` at the start of the next iteration.

`PathCollector` turns the resulting linked lists of flow into objects of struct type `Path`. The rooms belonging to each path, `p`, are stored in a slice in the `p.Room` field. The paths themselves are collected into a slice and ordered by length for ease of assigning the ants.

Future iterations of the path search revise and augment the flow, as described above. When no more paths can be found without breaking the capacity constraints or increasing the number of turns, the slice of paths is returned and used by `PrintTurns` to output the result.

To summarise `PathFinder`:

1. Set `numberOfTurns` to the maximum possible for any nest: `len(nest.Rooms) + ants - 2` (the number of rooms plus the number of ants minus two).
2. Begin loop.
3. Reset `Predecessor` field of all rooms to `nil`.
4. BFS.
5. If `nest.End` has no predecessor, break.
6. Update flow.
7. If this didn't reduce the number of turns, break.
8. Update paths.
9. End loop.
10. Return paths and flow.

And if flow is zero, the `main` function reports that no paths were found.

## 4. Further Notes

The formula for the maximum possible number of turns comes from the fact that this would occur if the graph consisted of one line of all nodes from `start` to `end`. Since the ants are already in `start`, we can subtract one from the number of turns it will take them to move through all the nodes. Since the first ant doesn't have to wait any turns, we can subtract another one, making a total of two. Consider, for example, the simplest case, where the nest consists of just two rooms, `start` and `end` and there is only one ant. Then `len(nest.Rooms) = 2` and `ants = 1`, so the number of turns is `2 + 1 - 2 = 1`.

In general, the number of turns taken will be the length of the path with the largest number of ants (which will be the first and shortest path, according to our way of assigning ants), plus the number of ants minus two.

Note that `u.Flow`, for a room `u`, is of type [`map[*Room]bool`](lem/structs.go). This booleam value is not quite what is meant by flow in the formal definition of a flow network. Rather, it's been streamlined to suit our case of unit capacity everywhere.

More properly, in any flow network (directed or undirected), if `f(u, v)` is the flow from a node `u` to another node `v`, then `f(v, u) = -f(u, v)`. This means that, for a directed graph with unit capacity `c(u, v) = 1` on an edge `(u, v)`, if we send flow along that edge, setting `f(u, v) = 1`, the residual capacity `cf` changes thus:

```
cf(u, v) = c(u, v) - f(u, v) = 1 - 1 = 0,
cf(v, u) = c(v, u) - f(v, u) = 0 - (-1) = 1,
```

representing the possibility, on a future attempt to find a path, to send flow in the opposite direction--in other words, to reverse our decision to send flow from `u` to `v`.

On the other hand, for our undirected graph, the capacity is one in both directions: `c(u, v) = c(v, u) = 1`, so if we send flow from `u` to `v`, the residual capacity becomes:

```
cf(u, v) = c(u, v) - f(u, v) = 1 - 1 = 0,
cf(v, u) = c(v, u) - f(v, u) = 1 - (-1) = 2,
```

which represents the possibility now to reverse our decision, cancelling out the flow from  `u` to `v` and then to still have the ability to send flow from `v` to `u`. However, since any path must send flow from `start` to an adjacent node (and likewise to `end` along an edge from one of its neighbours), and since these "forward" directions must have unit residual capacity, `1` is the "bottleneck" value for any path, and that full residual capacity of `2` on in a reverse direction can never be used. Because of this, our program uses a simplified definition of flow that only takes values of `0` or `1` and never `-1`.

## 5. CURIOSITIES

Depending on the network and number of ants, there may exist optimal solutions with fewer-than-maximal paths. The audit answer for [example05](nests/audit_examples/example05) is such a case. The number of ants is small enough to achieve the smallest number of turns with only three paths. However, as the number of ants is increased, eventually these three tunnels require more turns than our maximal solution of four paths. Thus, with nine ants, both solutions take eight turns, but, with 99 ants, ours takes 30 turns, while theirs takes 38.

More importantly, for the task of minimising the number of turns, consider [nests/sneaky_examples/few.txt](nests/sneaky_examples/few.txt). Here the maximum flow, consisting of two paths, actually takes more turns than the single shortest path when there are less than four ants, and only outperforms that short path when the number of ants is greater than five. This is why we need to make sure, after each BFS, that the new set of paths doesn't add to the number of turns taken.

More subtly, while our program gives a solution with the smallest number of turns, it can happen that other, shorter paths are available for the first few ants, permitting a solution with just as few turns, but even fewer individual ant-moves. This is the case in [example01](nests/audit_examples/example01), where the first ant to go to `h` can take one of the shorter paths, `start-h-n-e-end` or `start-n-m-end`, without blocking ants coming via `0` or `t`, provided all other ants follow the three longer paths of the maximum flow.

## 6. BIBLIOGRAPHY

[^0]: 01-Edu: Public Repo [ https://github.com/01-edu/public/tree/master/subjects/lem-in ]. Accessed Jan. 1, 2023.

[^D]: Dawson J: Lem-in: Finding all the paths and deciding which are worth it [ https://medium.com/@jamierobertdawson/lem-in-finding-all-the-paths-and-deciding-which-are-worth-it-2503dffb893 ]. Nov. 19, 2019. Accessed Jan. 1, 2023.

[^S]: Schroeder J, Guedes AP, Duarte EP: Computing the Minimum Cut and Maximum Flow of Undirected Graphs
[ https://www.inf.ufpr.br/pos/techreport/RT_DINF003_2004.pdf ]. RT-DINF 003/2004. Accessed Jan. 1. 2023.

[^W]: Wikipedia: Edmonds-Karp algorithm [ https://en.wikipedia.org/wiki/Edmonds%E2%80%93Karp_algorithm ]. Apr. 14, 2022. Accessed Jan. 1. 2023.
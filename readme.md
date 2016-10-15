## Suffix Tree

Suffix Trees are tree structures with a few types of elements.

### DataSource

A data source provides the sequence of numbers used in the creation of the suffix tree.  Each suffix is represented in the tree, starting at the root of the tree.

Non-numeric data types need to be converted to numeric values to be represented in a suffix tree.

### Node

There are three types of Nodes:

1. root -- the root of the tree
2. internal -- has an incoming edge, a suffix link to the next suffix, and at least two outgoing edges.
3. leaf -- there is one leaf associated with each suffix, it has an incoming edge, and no outgoing edges.

### Edge

Selects a sequence of values from a data source, referenced
as a outgoing edge of one Node and an incoming edge of another Node.

### Suffix Tree Construction

Each value added to the tree is added in constant time, allowing a tree size N to be constructed in O(n) time.

### Suffix Tree Queries

Suffix trees provide constant time responses to queries showing the location of an arbitrary sequence of values.
First the tree is traversed down the sequence of values, then the subtree is traversed (or precalculated) to
show the location of each value in the original sequence.


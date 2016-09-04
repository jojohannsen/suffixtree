## Suffix Tree

Suffix Trees are tree structures with a few types of elements.

### DataSource

A data source provides the sequence of numbers used in the creation of the suffix tree.  Each suffix is represented in the tree, starting at the root of the tree.

The values can be any data type, but are always mapped to a numeric value.

### Node

There are three types of Nodes:

1. root -- the root of the tree
2. internal -- has an incoming edge, a suffix link to the next suffix, and at least two outgoing edges.
3. leaf -- there is one leaf associated with each suffix, it has an incoming edge, and no outgoing edges.

### Edge

Selects a sequence of values from a data source, referenced
as a outgoing edge of one Node and an incoming edge of another Node.




## Suffix Tree

Suffix Tree are created from a few elements.

### Node

There are three types of Nodes:

1. root -- the top of the tree, only has outgoing edges
2. internal -- has one incoming edge, as well as outgoing edges.  During tree construction also has suffix links.
3. leaf -- has one incoming edge and a suffix offset

### Edge

Selects a sequence of values from a data source, referenced
as a outgoing edge of one Node and an incoming edge of another Node.

### Data Source



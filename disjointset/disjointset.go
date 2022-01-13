package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

// TODO: implement a type that satisfies the DisjointSet interface.
type MyDisjointSet map[int]int

var rank = make(map[int]int)

func (set MyDisjointSet) FindSet(number_key int) int {
	val, ok := set[number_key]
	if !ok {
		set[number_key] = number_key
		rank[number_key] = 0
		return number_key
	}
	if number_key == val {
		return number_key
	}
	// Path compression heuristic
	set[number_key] = set.FindSet(set[number_key])
	return set[number_key]
}

func (set MyDisjointSet) UnionSet(a int, b int) int {
	// Union by rank heuristic
	a = set.FindSet(a)
	b = set.FindSet(b)
	if a != b {
		if rank[a] < rank[b] {
			temp := a
			a = b
			b = temp
		}
		set[b] = a
		if rank[a] == rank[b] {
			rank[a]++
		}
	}
	return a
	/* number_key = set.FindSet(number_key)
	number_key2 = set.FindSet(number_key2)
	if number_key != number_key2 {
		set[number_key2] = number_key
	}
	return number_key */
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	set := make(MyDisjointSet)
	return set
}

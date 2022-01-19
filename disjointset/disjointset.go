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
type MyDisjointSet struct {
	set  map[int]int
	rank map[int]int
}

func (set MyDisjointSet) FindSet(number_key int) int {
	val, ok := set.set[number_key]
	if !ok {
		set.set[number_key] = number_key
		set.rank[number_key] = 0
		return number_key
	}
	if number_key == val {
		return number_key
	}
	// Path compression heuristic
	set.set[number_key] = set.FindSet(set.set[number_key])
	return set.set[number_key]
}

func (set MyDisjointSet) UnionSet(a int, b int) int {
	// Union by rank heuristic
	a = set.FindSet(a)
	b = set.FindSet(b)
	if a != b {
		if set.rank[a] < set.rank[b] {
			temp := a
			a = b
			b = temp
		}
		set.set[b] = a
		if set.rank[a] == set.rank[b] {
			set.rank[a]++
		}
	}
	return a
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	var set MyDisjointSet
	set.set = make(map[int]int)
	set.rank = make(map[int]int)
	return set
}

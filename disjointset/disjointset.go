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

func (set MyDisjointSet) FindSet(number_key int) int {
	if number_key == set[number_key] {
		return number_key
	}
	set[number_key] = set.FindSet(set[number_key])
	return set[number_key]
}

func (set MyDisjointSet) UnionSet(number_key int, number_key2 int) int {
	number_key = set.FindSet(number_key)
	number_key2 = set.FindSet(number_key2)
	if number_key != number_key2 {
		set[number_key2] = number_key
	}
	return number_key
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	// Insert all the elements in the set
	set := make(MyDisjointSet)
	for i := -1000 * 1000; i < 1000*1000; i++ {
		set[i] = i
	}
	return set
}

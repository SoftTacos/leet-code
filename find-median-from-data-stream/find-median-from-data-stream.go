package findmedianfromdatastream

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/*
Notes:
- could just toss everything into a slice(golang equivalent of a vector), and then sort it when FindMedian is called
  that would take O(nlog(n)) time, but would use O()
*/

// type MedianFinder struct {
// 	values []int
// }

// func Constructor() MedianFinder {
// 	return MedianFinder{}
// }

// func (mf *MedianFinder) AddNum(num int) {
// 	mf.values = append(mf.values, num)
// }

// func (mf *MedianFinder) FindMedian() (median float64) {
// 	sort.Slice(mf.values, func(i, j int) bool {
// 		return mf.values[i] < mf.values[j]
// 	})
// 	if len(mf.values)%2 == 0 {
// 		median = (float64(mf.values[(len(mf.values)-1)/2]) + float64(mf.values[(len(mf.values))/2])) / 2
// 	} else {
// 		median = float64(mf.values[len(mf.values)/2])
// 	}
// 	return
// }

// unbalanced BST
// takes longer to lookup the median if it isn't balanced
// however there seem to be more insertions than lookups
type Node struct {
	Val int
	// Count   int
	Less    *Node
	Greater *Node
}

func (n *Node) AddNum(num int) (newNode *Node) {
	// if n is nil, then we need to
	if n == nil {
		n = &Node{
			Val: num,
		}
	} else if num < n.Val {
		n.Less = n.Less.AddNum(num)
	} else {
		n.Greater = n.Greater.AddNum(num)
	}
	return n
}

// in order traversal, aka left node, current node, right node
func (n *Node) FindMedian(totalNodes int, count *int, median *float64) {

	// traverse left
	if n.Less != nil {
		n.Less.FindMedian(totalNodes, count, median)
	}

	// traverse self
	(*count)++
	// if this is a median node, then increment the median
	// even median node case
	if totalNodes%2 == 0 && ((totalNodes)/2 == *count || (totalNodes/2)+1 == *count) {
		*median = *median + (float64(n.Val) / 2)

		// odd median node case
	} else if (totalNodes+1)/2 == *count {
		*median = float64(n.Val)
	}

	// traverse right
	if n.Greater != nil {
		n.Greater.FindMedian(totalNodes, count, median)
	}
}

type MedianFinder struct {
	root       *Node
	totalNodes int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.totalNodes == 0 {
		mf.root = &Node{
			Val: num,
		}
	} else {
		mf.root.AddNum(num)
	}
	mf.totalNodes++
}

func (mf *MedianFinder) FindMedian() (median float64) {
	var count int
	mf.root.FindMedian(mf.totalNodes, &count, &median)
	return
}

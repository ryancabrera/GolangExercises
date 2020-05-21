//Based off of heap tutorial by Gayle Laakmann McDowell & "Golang By Example"
// https://www.youtube.com/watch?v=t0Cq6tVNRBA
// https://golangbyexample.com/heap-in-golang/
// https://blog.golang.org/slices

package heap

type MinIntHeap struct {
	capacity int
	size     int
	items    []int
}

//Operations to get Indexes of nodes
func (m *MinIntHeap) getLeftChildIndex(parentIndex int) int  { return 2*parentIndex + 1 }
func (m *MinIntHeap) getRightChildIndex(parentIndex int) int { return 2*parentIndex + 2 }
func (m *MinIntHeap) getParentIndex(childIndex int) int      { return (childIndex - 1) / 2 }

//Operations to check for existence of nodes
func (m *MinIntHeap) hasLeftChildIndex(index int) bool  { return m.getLeftChildIndex(index) < m.size }
func (m *MinIntHeap) hasRightChildIndex(index int) bool { return m.getRightChildIndex(index) < m.size }
func (m *MinIntHeap) hasParent(index int) bool          { return m.getParentIndex(index) >= 0 }

//Operations to get the Node Values
func (m *MinIntHeap) leftChild(index int) int  { return m.items[m.getLeftChildIndex(index)] }
func (m *MinIntHeap) rightChild(index int) int { return m.items[m.getRightChildIndex(index)] }
func (m *MinIntHeap) parent(index int) int     { return m.items[m.getParentIndex(index)] }

//Swap values at index
func (m *MinIntHeap) swap(indexOne int, indexTwo int) {
	m.items[indexOne], m.items[indexTwo] = m.items[indexTwo], m.items[indexOne]
}

//WIP TODO: Use slice functions to double capaciy of array
func (m *MinIntHeap) ensureExtraCapacity() {
	if m.size == m.capacity {
		//I'll have to take copy the values of the existing array into a new array, then point the old array to the new array, then clean up the old array
	}
}

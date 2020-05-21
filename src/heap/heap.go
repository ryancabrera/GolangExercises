//Based off of heap tutorial by Gayle Laakmann McDowell & "Golang By Example"
// https://www.youtube.com/watch?v=t0Cq6tVNRBA
// https://golangbyexample.com/heap-in-golang/
// https://blog.golang.org/slices

package heap

import "fmt"

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

//double capacity of array
func (m *MinIntHeap) ensureExtraCapacity() {
	if m.size == m.capacity {
		indexBeforeDoubling := len(m.items)
		m.items = append(m.items, make([]int, indexBeforeDoubling)...)
		m.items = m.items[:indexBeforeDoubling]
		m.capacity *= 2
	}
}

//Swap values at index
func (m *MinIntHeap) peek() int {
	if m.size == 0 {
		panic("Illegal State, Size is zero?")
	}
	return m.items[0]
}

func (m *MinIntHeap) poll() int {
	item := m.items[0]
	m.items[0] = m.items[m.size-1]
	m.size -= 1
	m.heapifyDown()
	return item
}

func (m *MinIntHeap) add(item int) {
	m.ensureExtraCapacity()
	m.items[m.size] = item
	m.size += 1
	m.heapifyUp()
}

func (m *MinIntHeap) heapifyUp() {
	index := m.size - 1
	for m.hasParent(index) && m.parent(index) > m.items[index] {
		m.swap(m.getParentIndex(index), index)
		index = m.getParentIndex(index)
	}
}

func (m *MinIntHeap) heapifyDown() {
	index := 0
	for m.hasLeftChildIndex(index) {
		smallerChildIndex := m.getLeftChildIndex(index)
		if m.hasRightChildIndex(index) && m.rightChild(index) < m.leftChild(index) {
			smallerChildIndex = m.getRightChildIndex(index)
		}

		if m.items[index] < m.items[smallerChildIndex] {
			break
		} else {
			m.swap(index, smallerChildIndex)
		}
		index = smallerChildIndex
	}
}

func Main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	heap := MinIntHeap{
		capacity: 2,
		size:     2,
		items:    inputArray,
	}

	heap.heapifyDown()

	for _, number := range heap.items {
		fmt.Println("Number is: ", number)
	}
}

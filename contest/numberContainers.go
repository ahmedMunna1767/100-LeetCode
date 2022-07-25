package contest

import (
	"sort"
)

type NumberContainers struct {
	container map[int][]int
}

func Constructor() NumberContainers {
	containerMap := make(map[int][]int)
	return NumberContainers{container: containerMap}
}

func (this *NumberContainers) Change(index int, number int) {

	for _, v := range this.container {
		for i := 0; i < len(v); i++ {
			if v[i] == index {
				v = append(v[:i], v[i+1:]...)
			}
		}
	}

	_, exists := this.container[number]
	if exists {
		this.container[number] = append(this.container[number], index)
	} else {
		set := make([]int, 0)
		set = append(set, index)
		this.container[number] = set
	}
}

func (this *NumberContainers) Find(number int) int {
	set, exists := this.container[number]
	if !exists {
		return -1
	}
	sort.Ints(set)

	return set[0]
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

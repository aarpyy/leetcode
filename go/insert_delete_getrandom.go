package leetcode

import (
	"math/rand"
)

type RandomizedSet struct {
	set  map[int]bool
	keys []int
}

func Constructor() RandomizedSet {
	randomizedSet := RandomizedSet{
		set:  make(map[int]bool, 10),
		keys: make([]int, 0, 10),
	}
	return randomizedSet
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}
	this.set[val] = true
	this.keys = append(this.keys, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.set[val]
	if !ok {
		return false
	}
	delete(this.set, val)
	for index, key := range this.keys {
		if key == val {
			this.keys = append(this.keys[:index], this.keys[index+1:]...)
			break
		}
	}
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.keys))
	return this.keys[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

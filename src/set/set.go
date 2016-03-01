package set

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}



/*
type Set_plus interface {
	IsSuperSet(other *HashSet) bool
	Union(other *HashSet) (union *HashSet)
	Intersect (other *HashSet) (intersect *HashSet)
	Difference(other *HashSet) (difference *HashSet)
	SymmetricDifference(other *HashSet) (symmetricdifference *HashSet)

}
*/


func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

func (set *HashSet) Len() int {
	return len(set.m)
}

func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// 此方法用于生成快照。
// This method is used to generate a snapshot.
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0

	//下面的语句中 key 用来获取map set 中的interface{} 对象 ，initialLen用来记录初始时set的长度，即set中的元素个数，snapshot 是一个接口类型的数组（切片（slice）） ，actualLen 是一个初始的索引标志。
	// The following statement in key used to obtain the map and set the object interface{}, used to record the initial set when the length of the initialLen, namely the number of elements in the set, the snapshot is an interface type array (section (slice), actualLen is of an initial index marker.
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

func (set *HashSet) IsSuperSet(other *HashSet) bool {
	if other == nil {
		return false
	}
	setLen := set.Len()
	otherLen := other.Len()
	if setLen == 0 || setLen == otherLen {
		return false
	}
	for _, v := range other.Elements() {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set *HashSet) Union(other *HashSet) (union *HashSet){
	if other == nil {
		return nil
	}
	if set.Same(other){
		return other
	}

	union=NewHashSet()

	for _,v:=range set.Elements() {
		err:=union.Add(v)
		if !err {
			fmt.Printf("Add %v Error!\n",v)
		}
	}

	for _,v:=range other.Elements() {
		if !union.Contains(v){
			union.Add(v)
		}
	}
	return union
}

func (set *HashSet) Intersect (other *HashSet) (intersect *HashSet) {
	if other == nil {
		return nil
	}
	intersect = NewHashSet()
	if set.Same(other){
		return intersect
	}

	for _,v := range set.Elements() {
		if other.Contains(v) {
			intersect.Add(v)
		}
	}

	return intersect
}

func (set *HashSet) Difference(other *HashSet) (difference *HashSet) {
	difference = NewHashSet()
	if other == nil {
		return nil
	}

	for _,v := range other.Elements() {
		if !set.Contains(v){
			difference.Add(v)
		}
	}

	return
}

func (set *HashSet)SymmetricDifference(other *HashSet) (symmetricdifference *HashSet) {
	symmetricdifference= NewHashSet()
	if other == nil {
		return nil
	}

	for _,v := range other.Elements() {
		if !set.Contains(v){
			symmetricdifference.Add(v)
		}
	}

	for _,v := range set.Elements() {
		if !other.Contains(v) {
			symmetricdifference.Add(v)
		}
	}
	return
}


type Set interface  {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other *HashSet) bool
	String() string
	Elements() []interface{}
}


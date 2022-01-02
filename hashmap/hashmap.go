package hashmap

const (
	md = 2048
)

type node struct {
	key  int
	val  int
	next *node
}

type MyHashMap struct {
	arr [md]*node
}

func Constructor() MyHashMap {
	return MyHashMap{
		arr: [md]*node{},
	}
}

func (mp *MyHashMap) Put(key, val int) {
	i := key % md
	p := mp.arr[i]

	for p != nil {
		if p.key == key {
			p.val = val
			return
		} else {
			p = p.next
		}
	}

	x := &node{key: key, val: val, next: mp.arr[key%md]}
	mp.arr[i] = x
}

func (mp *MyHashMap) Get(key int) int {
	i := key % md
	p := mp.arr[i]

	for p != nil {
		if p.key == key {
			return p.val
		}

		p = p.next
	}

	return -1
}

func (mp *MyHashMap) Remove(key int) {
	i := key % md
	curr := mp.arr[i]

	var prev *node

	for curr != nil {
		if curr.key == key {
			if prev == nil {
				mp.arr[i] = curr.next
			} else {
				prev.next = curr.next
			}
			break
		} else {
			prev = curr
			curr = curr.next
		}
	}

}

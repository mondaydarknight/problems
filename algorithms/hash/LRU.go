type Node struct {
    Prev, Next *Node
    Key, Val int
}

type LRUCache struct {
    capacity int
    hash map[int]*Node
    head, tail *Node
}

func Constructor(capacity int) LRUCache {
    head, tail := &Node{}, &Node{}
    head.Next, tail.Next = tail, head
    return LRUCache{capacity, make(map[int]*Node), head, tail}
}

func (this *LRUCache) Get(key int) int {
    if n, ok := this.hash[key]; ok {
        this.remove(n)
        this.add(n)
        return n.Val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.hash[key]; ok {
        this.remove(node)
    }
    if len(this.hash) == this.capacity {
        this.remove(this.tail.Prev)
    }
    this.add(&Node{Key: key, Val: value})
}

func (this *LRUCache) add(node *Node) {
    this.hash[node.Key] = node
    this.head.Next, node.Prev, node.Next, this.head.Next.Prev = node, this.head, this.head.Next, node
}

func (this *LRUCache) remove(node *Node) {
    delete(this.hash, node.Key)
    node.Prev.Next, node.Next.Prev = node.Next, node.Prev
}

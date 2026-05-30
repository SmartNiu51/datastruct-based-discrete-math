package container

// ValueIterator 表示可遍历单一类型元素序列的迭代器。
//
// Iter 返回一个 yield 函数，调用方可使用 range-over-func 语法：
//
//	for v := range it.Iter() {
//	    // 处理 v
//	}
//
// Items 一次性收集所有元素到切片中。对于大数据集，优先使用 Iter 以避免内存拷贝。
type ValueIterator[T any] interface {
	// Iter 返回一个符合 Go range-over-func 约定的迭代函数。
	Iter() (yield func(T) bool)
	// Items 返回容器中所有元素的切片副本。
	Items() []T
}

// KeyValueIterator 表示可遍历键值对的迭代器。
//
// Iter 返回的 yield 函数接收 key 和 value 两个参数：
//
//	for k, v := range it.Iter() {
//	    // 处理 k, v
//	}
type KeyValueIterator[K, V any] interface {
	// Iter 返回一个符合 Go range-over-func 约定的键值对迭代函数。
	Iter() (yield func(K, V) bool)
	// Items 返回所有键值对组成的切片。
	Items() []struct {
		Key   K
		Value V
	}
}

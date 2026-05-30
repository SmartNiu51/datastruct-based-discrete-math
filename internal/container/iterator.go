package container

// ValueIterator 表示可遍历单一类型元素序列的迭代器。
//
// 它提供两种遍历方式：
//   - Iter：惰性遍历，适合大数据集
//   - Values：一次性收集，适合需要随机访问的小数据集
//
// 典型用法：
//
//	for v := range it.Iter() {
//	    // 处理 v
//	}
type ValueIterator[T any] interface {
	// Iter 返回一个符合 Go range-over-func 约定的迭代函数。
	// 每次迭代产出序列中的一个元素，遍历顺序由实现决定。
	Iter() (yield func(T) bool)

	// Values 返回容器中所有元素的切片副本。
	// 返回的切片与容器内部数据无关，修改不影响原容器。
	// 对于大数据集，优先使用 Iter 以避免内存拷贝。
	Values() []T
}

// IndexValueIterator 表示遍历时可同时获取索引和值的迭代器。
//
// 它和 ValueIterator 的区别在于 Iter 的 yield 函数额外提供位置索引。
// 适用于列表、数组等有明确位置概念的有序容器。
//
// 典型用法：
//
//	for i, v := range it.Iter() {
//	    // 处理索引 i 和值 v
//	}
type IndexValueIterator[T any] interface {
	// Iter 返回一个符合 Go range-over-func 约定的索引-值迭代函数。
	// 每次迭代产出 (索引, 值) 对，索引从 0 开始递增。
	Iter() (yield func(int, T) bool)

	// Values 返回容器中所有元素的切片副本，按迭代顺序排列。
	// 返回的切片不含索引信息，仅包含值。
	Values() []T
}

// KeyValueIterator 表示可遍历键值对的迭代器。
//
// 适用于映射、字典等关联容器。每次迭代产出一个键值对。
//
// 典型用法：
//
//	for k, v := range it.Iter() {
//	    // 处理键 k 和值 v
//	}
type KeyValueIterator[K, V any] interface {
	// Iter 返回一个符合 Go range-over-func 约定的键值对迭代函数。
	// 每次迭代产出 (键, 值) 对。遍历顺序由实现决定，不保证有序。
	Iter() (yield func(K, V) bool)

	// Pairs 返回所有键值对组成的切片副本。
	// 返回的切片与容器内部数据无关。
	Pairs() []struct {
		Key   K
		Value V
	}
}

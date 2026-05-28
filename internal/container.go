// Package basic 提供所有数据结构的基础接口定义。
//
// 本包定义了贯穿整个库的原子接口（Sized、Clearable、Membership 等）
// 及其复合接口（Container、MutableContainer）。
// 所有上层数据结构接口均通过嵌入本包的接口来声明其行为契约。
package basic

// Sized 表示可查询元素数量的容器。
type Sized interface {
	// Sized 返回容器中元素的数量。
	Sized() int
	// Empty 返回容器是否为空。
	Empty() bool
}

// Clearable 表示可清空全部元素的容器。
type Clearable interface {
	// Clear 清空容器中的所有元素。
	Clear()
}

// Membership 表示可判断元素是否存在的容器。
type Membership[T any] interface {
	// Contains 判断元素 v 是否在容器中。
	Contains(T) bool
}

// Container 是可查询大小、可清空、可判断成员的基础容器接口。
//
// 它是 Sized、Clearable 和 Membership 的并集，
// 被所有具体数据结构接口所嵌入。
type Container[T any] interface {
	Sized
	Clearable
	Membership[T]
}

// MutableContainer 是可增删元素的可变容器接口。
//
// 它嵌入 Container 的全部只读行为，并额外提供 Add 和 Remove。
type MutableContainer[T any] interface {
	Container[T]
	// Add 向容器中添加元素。
	Add(T)
	// Remove 从容器中移除元素，存在并移除成功时返回 true。
	Remove(T) bool
}

// SetOperations 表示可执行集合代数运算的类型。
//
// 所有运算返回一个新的实例，不修改原容器。
type SetOperations[T any] interface {
	// Union 返回与 other 的并集。
	Union(other T) T
	// Intersect 返回与 other 的交集。
	Intersect(other T) T
	// Difference 返回与 other 的差集（本容器有而 other 无的元素）。
	Difference(other T) T
}

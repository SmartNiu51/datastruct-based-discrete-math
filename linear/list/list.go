// Package list 提供 List[T] 接口及其辅助函数。
//
// List 表示支持随机访问的线性表，可任意位置读写、增删元素。
// 不同实现（动态数组、链表、跳表等）均满足此接口。
package list

import (
	container "github.com/SmartNiu51/datastruct-based-discrete-math/internal/container"
)

// List 表示支持随机访问的线性表。
//
// 它组合了三个基础能力：
//   - Container：容量感知、可清空、成员判断
//   - IndexValueIterator：带索引的遍历
//
// 并额外提供任意位置的读写、插入、删除和查找。
type List[T any] interface {
	container.Container[T]
	container.IndexValueIterator[T]

	// Get 返回索引 i 处的元素。越界返回零值 + false。
	Get(index int) (T, bool)

	// Set 设置索引 i 处的值为 val。越界返回 false。
	Set(index int, val T) bool

	// Insert 在索引 i 处插入 val，后续元素后移。
	// i ∈ [0, Size()]，i == Size() 时等价于追加到末尾。越界返回 false。
	Insert(index int, val T) bool

	// RemoveAt 移除并返回索引 i 处的元素。越界返回零值 + false。
	RemoveAt(index int) (T, bool)

	// FirstIndexOf 返回 val 首次出现的索引。未找到返回 -1。
	FirstIndexOf(val T) int

	// LastIndexOf 返回 val 最后一次出现的索引。未找到返回 -1。
	LastIndexOf(val T) int
}

// Append 在列表末尾批量追加元素。
func Append[T any](list List[T], values ...T) {
	// TODO: 实现批量追加
}

// Swap 交换索引 i 和 j 处的元素。越界返回 false。
func Swap[T any](list List[T], i, j int) bool {
	// TODO: 实现交换逻辑
	return false
}

// Reverse 原地反转列表。
func Reverse[T any](list List[T]) {
	// TODO: 实现反转逻辑
}

// Sort 按 compare 函数原地排序列表。
type cmp[T any] = container.CmpFunc[T]

func Sort[T any](list List[T], compare cmp[T]) {
	// TODO: 实现排序逻辑
}

// RemoveRange 移除 [from, to) 范围内的元素，返回实际移除的数量。
func RemoveRange[T any](list List[T], from, to int) int {
	// TODO: 实现区间删除
	return 0
}

// SubList 返回 [from, to) 范围内的新列表。越界返回 nil, false。
func SubList[T any](list List[T], from, to int) (List[T], bool) {
	// TODO: 实现子列表
	return nil, false
}

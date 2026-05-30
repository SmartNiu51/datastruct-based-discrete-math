package list

import (
	container "github.com/SmartNiu51/datastruct-based-discrete-math/internal/container"
)

// Append 在列表末尾批量追加元素。
func Append[T any](list List[T], values ...T) {
	// TODO: 实现批量追加
	panic("unimplemented")
}

// Swap 交换索引 i 和 j 处的元素。越界返回 false。
func Swap[T any](list List[T], i, j int) bool {
	// TODO: 实现交换逻辑
	panic("unimplemented")
}

// Reverse 原地反转列表。
func Reverse[T any](list List[T]) {
	// TODO: 实现反转逻辑
	panic("unimplemented")
}

// Sort 按 compare 函数原地排序列表。
type cmp[T any] = container.CmpFunc[T]

func Sort[T any](list List[T], compare cmp[T]) {
	// TODO: 实现排序逻辑
	panic("unimplemented")
}

// RemoveRange 移除 [from, to) 范围内的元素，返回实际移除的数量。
func RemoveRange[T any](list List[T], from, to int) int {
	// TODO: 实现区间删除
	panic("unimplemented")
}

// SubList 返回 [from, to) 范围内的新列表。越界返回 nil, false。
func SubList[T any](list List[T], from, to int) (List[T], bool) {
	// TODO: 实现子列表
	panic("unimplemented")
}

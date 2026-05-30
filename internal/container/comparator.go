package container

// Comparator 定义两个同类型元素的比较语义。
//
// Compare 的返回值约定：
//
//	负值  → a < b
//	零    → a == b
//	正值  → a > b
//
// 典型实现：
//
//	type IntComparator struct{}
//	func (IntComparator) Compare(a, b int) int { return a - b }
type Comparator[T any] interface {
	Compare(a, b T) int
}

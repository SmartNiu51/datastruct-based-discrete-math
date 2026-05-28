# datastruct-based-discrete-math

基于离散数学的 Go 泛型数据结构与算法库。

## 设计理念

- **接口即契约** — 所有领域概念定义为 `interface`，外部不感知 `struct`
- **组合即继承** — 接口嵌入实现继承关系，实现层用具名组合
- **语义兼容性** — 仅当数学上满足 is-a 时才使用接口嵌入
- **算法与结构解耦** — 算法为独立 `algo/` 包，纯函数仅依赖接口

## 当前进度

Phase 1（进行中）— 基础数据结构与核心算法

- [ ] 原子接口层（`container/`）
- [ ] 线性结构（Stack / Queue / Deque / List / PriorityQueue）
- [ ] 集合与映射（Set / Map）
- [ ] 堆（BinaryHeap）
- [ ] 图论基础（Graph / Digraph / WeightedGraph / Tree / BinaryTree / BST）
- [ ] 核心算法（BFS / DFS / Dijkstra / 拓扑排序）

## 安装

```bash
go get github.com/SmartNiu51/datastruct-based-discrete-math
```

要求 Go >= 1.26。

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/SmartNiu51/datastruct-based-discrete-math/ds"
)

func main() {
    s := ds.NewStack[int]()
    s.Push(1)
    s.Push(2)
    s.Push(3)

    for !s.IsEmpty() {
        v, _ := s.Pop()
        fmt.Println(v)
    }
    // Output:
    // 3
    // 2
    // 1
}
```

## 包索引

| 包 | 内容 |
|----|------|
| `container` | 原子接口：`Iterable`, `Container`, `Collection`, `Indexable`, `Associative` |
| `ds` | 线性结构：`Stack`, `Queue`, `Deque`, `List`, `PriorityQueue` |
| `set` | 集合论：`Set`, `MultiSet`, `DisjointSet` |
| `map` | 映射：`Map`, `BiMap`, `MultiMap` |
| `heap` | 堆族：`Heap`, `BinaryHeap` |
| `graph` | 图论：`Graph`, `Digraph`, `WeightedGraph`, `Tree`, `BinaryTree`, `BST` |
| `algo` | 算法：BFS, DFS, Dijkstra, 拓扑排序 |

## 文档

| 文档 | 说明 |
|------|------|
| [架构设计](local/architecture_design.md) | 全局接口继承拓扑、目录结构、160+ 数据结构总览 |
| [第一阶段方案](local/phase1_plan.md) | 逐文件实施计划 |
| [开发规范](local/package_rules.md) | 接口化、组合、隔离、错误处理等架构铁律 |
| [测试规范](local/test_rules.md) | 单元/示例/集成测试写法与位置 |

## 许可证

MIT

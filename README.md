# [LeetCode](https://leetcode.com) 的 Go 解答

[![LeetCode 排名](https://img.shields.io/badge/aQuaYi-778-blue.svg)](https://leetcode.com/aQuaYi/)
[![codecov](https://codecov.io/gh/aQuaYi/LeetCode-in-Go/branch/master/graph/badge.svg)](https://codecov.io/gh/aQuaYi/LeetCode-in-Go)
[![Build Status](https://www.travis-ci.org/aQuaYi/LeetCode-in-Go.svg?branch=master)](https://www.travis-ci.org/aQuaYi/LeetCode-in-Go)

## 进度

> 统计规则：1.免费题，2.算法题，3.能用 Go 解答

|     |Easy|Medium|Hard|Total|
|:---:|:---:|:---:|:---:|:---:|
|**Accepted**|189|325|137|651|
|**Total**|193|331|141|665|

## 题解

|题号|题目|通过率|难度|收藏|
|:-:|:-|:-: | :-: | :-: |
|206|[Reverse Linked List](./Algorithms/0206.reverse-linked-list)|48%|Easy||

## helper

[helper](./helper) 会处理大部分琐碎的工作。

## notes

[notes](./notes) 记录了我答题过程中，对知识点的总结。

## kit

针对 LeetCode 中经常出现的以下数据结构，在 [kit](./kit) 中进行了定义，并添加了与 []int 相互转换的函数。利用 Go 1.9 新添加的 [type alias](https://github.com/golang/proposal/blob/master/design/18130-type-alias.md) 功能，易于添加单元测试。

- [Heap](./kit/Heap.go)
- [Interval](./kit/Interval.go)
- [ListNode](./kit/ListNode.go)
- [NestedInteger](./kit/NestedInteger.go)
- [PriorityQueue](./kit/PriorityQueue.go)
- [Queue](./kit/Queue.go)
- [Stack](./kit/Stack.go)
- [TreeNode](./kit/TreeNode.go)
- [Master](./kit/master.go)

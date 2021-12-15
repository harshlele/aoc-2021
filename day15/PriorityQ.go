package main

import (
	utils "aoc-2021/aoc-utils"
)

//WARNING - NEVER USE FOR LARGE QUEUES(ie part 2 lmao)
//with this, part 2 takes 1.5-2 hours to complete
type PriorityQ struct {
	Map map[string][]int
	min string
}

func (q *PriorityQ) Push(i, j int, val []int) {
	key := utils.PointKey([]int{i, j})
	if !q.isInMap(key) {
		q.Map[key] = val
		//update min
		if len(q.Map) > 1 {
			if val[1] < q.Map[q.min][1] {
				q.min = key
			}
		} else {
			q.min = key
		}
	}
}

func (q *PriorityQ) Remove(i, j int) []int {
	key := utils.PointKey([]int{i, j})
	if q.isInMap(key) {
		val := q.Map[key]
		delete(q.Map, key)
		if key == q.min {
			q.updateMin()
		}
		return val
	}
	return []int{-1, -1}
}

func (q *PriorityQ) UpdateAtKey(i, j int, newVal []int) {
	key := utils.PointKey([]int{i, j})
	if q.isInMap(key) {
		old := q.Map[key]
		q.Map[key] = newVal
		if key == q.min {
			if old[1] <= newVal[1] {
				q.updateMin()
			}
		}
	}

}

func (q *PriorityQ) updateMin() {
	min := -1
	currMinKey := ""
	for k := range q.Map {
		if min == -1 || q.Map[k][1] < min {
			min = q.Map[k][1]
			currMinKey = k
		}
	}
	q.min = currMinKey
}

func (q *PriorityQ) isInMap(key string) bool {
	_, ok := q.Map[key]
	if !ok {
		return false
	}
	return true
}

func (q *PriorityQ) MinVal() []int {
	return q.Map[q.min]
}

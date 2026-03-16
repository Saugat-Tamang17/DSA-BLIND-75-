package L_queue

func IsEmpty(q *Lqueue) bool {
	return q.Front == -1 || q.Front > q.Rear
}

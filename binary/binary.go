package binary

func Binary(list []int, item int) (int, error) {
	low := 0
	high := len(list)

	for low <= high {
		mid := low + high
		guess := list[mid]

		if guess == item {
			return mid, nil
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, &NotFoundError{"Item not found in list"}
}

type NotFoundError struct {
	msg string
}

func (e *NotFoundError) Error() string {
	return e.msg
}

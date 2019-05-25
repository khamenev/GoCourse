package statistic

/*
** 15.05.2019, Sergei Khamenev, Task-1
 */

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Функция подсчета суммы
func Sum(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}

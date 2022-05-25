package main

func sortMapValues(input map[int]string) (result []string) {
	var ar []int
	for i, val := range input {
		ar = append(ar, i)
		result = append(result, val)
	}
	length := len(ar)
	for i := length - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if ar[i] < ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return
}

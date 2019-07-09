package main

type val struct {
	val uint64
	key rune
}
type valW struct {
	val uint64
	key string
}

type vals []val
type valsW []valW

func sort(fr map[rune]uint64) vals {
	var i uint64
	out := make(vals, len(fr))
	for k, v := range fr {
		out[i].val = v
		out[i].key = k
		i++
	}
	qsort(out)
	return out
}

func sortW(fr map[string]uint64) valsW {
	var i uint64
	out := make(valsW, len(fr))
	for k, v := range fr {
		out[i].val = v
		out[i].key = k
		i++
	}
	qsortW(out)
	return out
}

func qsort(array vals) vals {
	var pivot uint64
	var i uint64

	if len(array) <= 1 {
		return array
	}
	for i = 1; i < uint64(len(array)); i++ {
		if array[i].val > array[0].val {
			pivot++
			array.swap(i, pivot)
		}
	}
	array.swap(0, pivot)
	qsort(array[0:pivot])
	qsort(array[pivot+1:])
	return array
}

func qsortW(array valsW) valsW {
	var pivot uint64
	var i uint64

	if len(array) <= 1 {
		return array
	}
	for i = 1; i < uint64(len(array)); i++ {
		if array[i].val > array[0].val {
			pivot++
			array.swapW(i, pivot)
		}
	}
	array.swapW(0, pivot)
	qsortW(array[0:pivot])
	qsortW(array[pivot+1:])
	return array
}

func (array vals) swap(i, j uint64) {
	var swp val

	swp = array[i]
	array[i] = array[j]
	array[j] = swp
}

func (array valsW) swapW(i, j uint64) {
	var swp valW

	swp = array[i]
	array[i] = array[j]
	array[j] = swp
}

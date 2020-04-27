package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(fn func(int) bool) Ints {

	result := Ints{}

	if len(ints) == 0 {
		return nil
	}
	for _, i := range ints {
		if fn(i) {
			result = append(result, i)
		}
	}

	return result
}

func (str Strings) Keep(fn func(string) bool) Strings {
	if len(str) == 0 {
		return nil
	}
	result := Strings{}
	for _, s := range str {
		if fn(s) {
			result = append(result, s)
		}
	}
	return result
}

func (ints Ints) Discard(fn func(int) bool) Ints {
	if len(ints) == 0 {
		return nil
	}

	result := Ints{}

	for _, i := range ints {
		if !fn(i) {
			result = append(result, i)
		}
	}
	return result

}

func (lists Lists) Keep(fn func([]int) bool) Lists {
	if len(lists) == 0 {
		return nil
	}
	result := Lists{}

	for _, l := range lists {
		if fn(l) {
			result = append(result, l)
		}
	}
	return result
}

package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(n int) int

func (l IntList) Foldl(f binFunc, x int) int {
	for _, v := range l {
		x = f(x, v)
	}
	return x
}

func (l IntList) Foldr(f binFunc, x int) int {
	end := len(l) - 1
	for i := range l {
		x = f(l[end-i], x)
	}
	return x
}

func (l IntList) Filter(f predFunc) IntList {
	out := l[:0]
	for _, v := range l {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func (l IntList) Length() int {
	return len(l)
}

func (l IntList) Map(f unaryFunc) IntList {
	for i, v := range l {
		l[i] = f(v)
	}
	return l
}

func (l IntList) Reverse() IntList {
	end := len(l) - 1 // index of final element
	for i := len(l)/2-1; i >= 0; i-- {
		l[i], l[end-i] = l[end-i], l[i]
	}
	return l
}

func (l IntList) Append(m IntList) IntList {
	return append(l, m...)
}

func (l IntList) Concat(args []IntList) IntList {
	for _, v := range args {
		l = append(l, v...)
	}
	return l
}

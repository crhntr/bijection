package bijection

func Make[xt, yt comparable]() Map[xt, yt] {
	return Map[xt, yt]{
		x2y: make(map[xt]yt),
		y2x: make(map[yt]xt),
	}
}

func MakeN[xt, yt comparable](n int) Map[xt, yt] {
	return Map[xt, yt]{
		x2y: make(map[xt]yt, n),
		y2x: make(map[yt]xt, n),
	}
}

type Map[XT, YT comparable] struct {
	x2y map[XT]YT
	y2x map[YT]XT
}

func (mp Map[XT, YT]) GetY(key XT) (YT, bool) {
	v, ok := mp.x2y[key]
	return v, ok
}

func (mp Map[XT, YT]) GetX(key YT) (XT, bool) {
	v, ok := mp.y2x[key]
	return v, ok
}

func (mp Map[XT, YT]) Set(key XT, value YT) {
	mp.x2y[key] = value
	mp.y2x[value] = key
}

func (mp Map[XT, YT]) DeleteX(key XT) {
	otherKey, ok := mp.x2y[key]
	if !ok {
		return
	}
	delete(mp.x2y, key)
	delete(mp.y2x, otherKey)
}

func (mp Map[XT, YT]) DeleteY(key YT) {
	otherKey, ok := mp.y2x[key]
	if !ok {
		return
	}
	delete(mp.y2x, key)
	delete(mp.x2y, otherKey)
}

func SwapRangeFunc[T1, T2 interface{}](fn func(T1, T2) bool) func(T2, T1) bool {
	return func(t2 T2, t1 T1) bool {
		return fn(t1, t2)
	}
}

func (mp Map[XT, YT]) Range(fn func(key XT, value YT) (continueIter bool)) {
	for key, value := range mp.x2y {
		continueIter := fn(key, value)
		if !continueIter {
			break
		}
	}
}

func (mp Map[XT, YT]) Inverse() Map[YT, XT] {
	inverted := MakeN[YT, XT](len(mp.x2y))

	mp.Range(func(key XT, value YT) (continueIter bool) {
		inverted.Set(value, key)
		return true
	})

	return inverted
}

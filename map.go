package bijection

type Map[A, B comparable] struct {
	ab map[A]B
	ba map[B]A
}

func (mp Map[A, B]) GetA(key B) (A, bool) {
	if mp.ab == nil {
		var v A
		return v, false
	}
	v, ok := mp.ba[key]
	return v, ok
}

func (mp *Map[A, B]) GetB(key A) (B, bool) {
	if mp.ab == nil {
		var v B
		return v, false
	}
	v, ok := mp.ab[key]
	return v, ok
}

func (mp *Map[A, B]) Add(key A, value B) {
	if mp.ab == nil {
		mp.ab, mp.ba = make(map[A]B), make(map[B]A)
	}
	mp.ab[key], mp.ba[value] = value, key
}

func (mp Map[A, B]) DeleteA(key A) {
	if mp.ab == nil {
		return
	}
	a := key
	b, ok := mp.ab[key]
	if !ok {
		return
	}
	delete(mp.ab, a)
	delete(mp.ba, b)
}

func (mp Map[A, B]) DeleteB(key B) {
	if mp.ab == nil {
		return
	}
	b := key
	a, ok := mp.ba[key]
	if !ok {
		return
	}
	delete(mp.ab, a)
	delete(mp.ba, b)
}

func SwapRangeFunc[T1, T2 comparable](fn func(T1, T2)) func(T2, T1) {
	return func(t2 T2, t1 T1) { fn(t1, t2) }
}

func (mp Map[A, B]) Range(fn func(key A, value B)) {
	for key, value := range mp.ab {
		fn(key, value)
	}
}

func (mp Map[A, B]) Inverse() Map[B, A] {
	var inverted Map[B, A]
	mp.Range(func(key A, value B) {
		inverted.Add(value, key)
	})
	return inverted
}

func (mp Map[A, B]) Len() int {
	if mp.ab == nil {
		return 0
	}
	return len(mp.ab)
}

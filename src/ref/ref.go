package ref_P

import gleam_P "example.com/todo/gleam"

type RefCell_t[T gleam_P.Type[T]] struct {
	state *state[T]
}

type state[T gleam_P.Type[T]] struct {
	value  T
	killed bool
}

func Cell[T gleam_P.Type[T]](val T) RefCell_t[T] {
	return RefCell_t[T]{&state[T]{value: val}}
}

func Get[T gleam_P.Type[T]](ref RefCell_t[T]) T {
	return ref.state.value
}

func TryGet[T gleam_P.Type[T]](ref RefCell_t[T]) gleam_P.Result_t[T, gleam_P.String_t] {
	if !ref.state.killed {
		return gleam_P.Ok_c[T, gleam_P.String_t]{ref.state.value}
	} else {
		return gleam_P.Error_c[T, gleam_P.String_t]{"Is this ok"}
	}
}

func Set[T gleam_P.Type[T]](ref RefCell_t[T], fun func(T) T) T {
	ref.state.value = fun(ref.state.value)
	return ref.state.value
}

func Kill[T gleam_P.Type[T]](ref RefCell_t[T]) {
	ref.state.killed = true
}

package longlived_grpc

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/constraints"
)

func ErrOr[T any](err error, a T) any {
	if err != nil {
		return err.Error()
	}

	return a
}

func IfAny(condition bool, a, b any) any {
	if condition {
		return a
	}
	return b
}

func IfError[T constraints.Ordered](err error, a, b T) T {
	if err != nil {
		return a
	}

	return b
}

func Max[T constraints.Ordered](s ...T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}

	m := s[0]
	for _, v := range s[1:] {
		if m < v {
			m = v
		}
	}
	return m
}

func Min[T constraints.Ordered](s ...T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}

	m := s[0]
	for _, v := range s[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func QueryInt(g *gin.Context, name string, min int) int {
	n := 0
	if q := g.Query(name); q != "" {
		n, _ = strconv.Atoi(q)
	}

	return Max(n, min)
}

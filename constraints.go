package goterators

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

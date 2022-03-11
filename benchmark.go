package testes

import (
	"github.com/skeptycal/types"
)

const (
	// scaling factor (in powers of 2)
	defaultMaxScalingFactor = 6
	maxScalingFactor        = 10
)

var (
	ValueOf = types.ValueOf
	Log     = types.Log
)

type (
	Any = types.Any

	AnyValue = types.AnyValue

	GetSetter interface {
		Get(key Any) (Any, error)
		Set(key Any, value Any) error
	}
)

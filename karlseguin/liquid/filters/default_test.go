package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
	"github.com/unix-world/go-modules/karlseguin/liquid/core"
)

func TestDefaultWithBuiltinValue(t *testing.T) {
	spec := gspec.New(t)
	filter := DefaultFactory(nil)
	spec.Expect(filter(nil, nil).(string)).ToEqual("")
}

func TestDefaultWithValueOnString(t *testing.T) {
	spec := gspec.New(t)
	filter := DefaultFactory([]core.Value{stringValue("d")})
	spec.Expect(filter("", nil).(string)).ToEqual("d")
}

func TestDefaultWithValueOnArray(t *testing.T) {
	spec := gspec.New(t)
	filter := DefaultFactory([]core.Value{stringValue("n/a")})
	spec.Expect(filter([]int{}, nil).(string)).ToEqual("n/a")
}

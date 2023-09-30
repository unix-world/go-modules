package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
	"github.com/unix-world/go-modules/karlseguin/liquid/core"
)

func TestRemovesFirstValueFromAString(t *testing.T) {
	spec := gspec.New(t)
	filter := RemoveFirstFactory([]core.Value{stringValue("foo")})
	spec.Expect(filter("foobarforfoo", nil).(string)).ToEqual("barforfoo")
}

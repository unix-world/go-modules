package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
	"github.com/unix-world/go-modules/karlseguin/liquid/core"
)

func TestPrependToAString(t *testing.T) {
	spec := gspec.New(t)
	filter := PrependFactory([]core.Value{stringValue("?!")})
	spec.Expect(filter("dbz", nil).(string)).ToEqual("?!dbz")
}

func TestPrependToBytes(t *testing.T) {
	spec := gspec.New(t)
	filter := PrependFactory([]core.Value{stringValue("boring")})
	spec.Expect(filter([]byte("so"), nil).(string)).ToEqual("boringso")
}

func TestPrependADynamicValue(t *testing.T) {
	spec := gspec.New(t)
	filter := PrependFactory([]core.Value{dynamicValue("local.currency")})
	data := map[string]interface{}{
		"local": map[string]string{
			"currency": "$",
		},
	}
	spec.Expect(filter("100", data).(string)).ToEqual("$100")
}

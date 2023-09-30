package filters

import (
	"testing"

	"github.com/karlseguin/gspec"
	"github.com/unix-world/go-modules/karlseguin/liquid/core"
)

func TestTruncateAString(t *testing.T) {
	spec := gspec.New(t)
	filter := TruncateFactory([]core.Value{intValue(5)})
	spec.Expect(filter("1234567", nil).(string)).ToEqual("12...")
}

func TestTruncateAShortString(t *testing.T) {
	spec := gspec.New(t)
	filter := TruncateFactory([]core.Value{intValue(100), stringValue("")})
	spec.Expect(filter("1234567", nil).(string)).ToEqual("1234567")
}

func TestTruncateAPerfectString(t *testing.T) {
	spec := gspec.New(t)
	filter := TruncateFactory([]core.Value{intValue(7), stringValue("")})
	spec.Expect(filter("1234567", nil).(string)).ToEqual("1234567")
}

func TestTruncateAnAlmostPerfectString(t *testing.T) {
	spec := gspec.New(t)
	filter := TruncateFactory([]core.Value{intValue(6), stringValue("")})
	spec.Expect(filter("1234567", nil).(string)).ToEqual("123456")
}

func TestTruncateAStringFromAFloat(t *testing.T) {
	spec := gspec.New(t)
	filter := TruncateFactory([]core.Value{floatValue(3.3), stringValue(".")})
	spec.Expect(filter("1234567", nil).(string)).ToEqual("12.")
}

func TestTruncateAStringFromAString(t *testing.T) {
	spec := gspec.New(t)
	filter := TruncateFactory([]core.Value{stringValue("4"), stringValue("")})
	spec.Expect(filter("1234567", nil).(string)).ToEqual("1234")
}

func TestTruncateAStringFromAnInvalidString(t *testing.T) {
	spec := gspec.New(t)
	filter := TruncateFactory([]core.Value{stringValue("abc"), stringValue("")})
	spec.Expect(filter("1234567", nil).(string)).ToEqual("1234567")
}

func TestTruncateAnInvalidValue(t *testing.T) {
	spec := gspec.New(t)
	filter := TruncateFactory([]core.Value{intValue(4)})
	spec.Expect(filter(555, nil).(string)).ToEqual("5...")
}

package filterhelper // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/processor/filterhelper"

import (
	"fmt"

	"github.com/spf13/cast"
	"go.opentelemetry.io/collector/model/pdata"
)

// NewAttributeValueRaw is used to convert the raw `value` from ActionKeyValue to the supported trace attribute values.
// If error different than nil the return value is invalid. Calling any functions on the invalid value will cause a panic.
func NewAttributeValueRaw(value interface{}) (pdata.AttributeValue, error) {
	switch val := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return pdata.NewAttributeValueInt(cast.ToInt64(val)), nil
	case float32, float64:
		return pdata.NewAttributeValueDouble(cast.ToFloat64(val)), nil
	case string:
		return pdata.NewAttributeValueString(val), nil
	case bool:
		return pdata.NewAttributeValueBool(val), nil
	default:
		return pdata.AttributeValue{}, fmt.Errorf("error unsupported value type \"%T\"", value)
	}
}

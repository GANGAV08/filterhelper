package filterhelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/model/pdata"
)

func TestHelper_AttributeValue(t *testing.T) {
	val, err := NewAttributeValueRaw(uint8(123))
	assert.Equal(t, pdata.NewAttributeValueInt(123), val)
	assert.NoError(t, err)

	val, err = NewAttributeValueRaw(uint16(123))
	assert.Equal(t, pdata.NewAttributeValueInt(123), val)
	assert.NoError(t, err)

	val, err = NewAttributeValueRaw(int8(123))
	assert.Equal(t, pdata.NewAttributeValueInt(123), val)
	assert.NoError(t, err)

	val, err = NewAttributeValueRaw(int16(123))
	assert.Equal(t, pdata.NewAttributeValueInt(123), val)
	assert.NoError(t, err)

	val, err = NewAttributeValueRaw(float32(234.129312))
	assert.Equal(t, pdata.NewAttributeValueDouble(float64(float32(234.129312))), val)
	assert.NoError(t, err)

	val, err = NewAttributeValueRaw(234.129312)
	assert.Equal(t, pdata.NewAttributeValueDouble(234.129312), val)
	assert.NoError(t, err)

	val, err = NewAttributeValueRaw(true)
	assert.Equal(t, pdata.NewAttributeValueBool(true), val)
	assert.NoError(t, err)

	val, err = NewAttributeValueRaw("bob the builder")
	assert.Equal(t, pdata.NewAttributeValueString("bob the builder"), val)
	assert.NoError(t, err)

	_, err = NewAttributeValueRaw(nil)
	assert.Error(t, err)

	_, err = NewAttributeValueRaw(t)
	assert.Error(t, err)
}

package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsumptionRequest_JSONSerialization(t *testing.T) {
	t.Run("should omit ConsumptionPercentage when nil", func(t *testing.T) {
		t.Parallel()
		req := ConsumptionRequest{
			CustomerConsented:     true,
			ConsumptionPercentage: nil,
			DeliveryStatus:        DELIVERED,
			RefundPreference:      DECLINE,
			SampleContentProvided: false,
		}

		data, err := json.Marshal(req)
		assert.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		assert.NoError(t, err)

		_, exists := result["consumptionPercentage"]
		assert.False(t, exists, "consumptionPercentage should be omitted")
	})

	t.Run("should include ConsumptionPercentage when set to 0", func(t *testing.T) {
		t.Parallel()
		zero := int32(0)
		req := ConsumptionRequest{
			CustomerConsented:     true,
			ConsumptionPercentage: &zero,
			DeliveryStatus:        DELIVERED,
			RefundPreference:      GRANT_FULL,
			SampleContentProvided: true,
		}

		data, err := json.Marshal(req)
		assert.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		assert.NoError(t, err)

		val, exists := result["consumptionPercentage"]
		assert.True(t, exists)
		assert.Equal(t, float64(0), val)
	})

	t.Run("should include ConsumptionPercentage when set to non-zero value", func(t *testing.T) {
		t.Parallel()
		fifty := int32(50)
		req := ConsumptionRequest{
			CustomerConsented:     true,
			ConsumptionPercentage: &fifty,
			DeliveryStatus:        DELIVERED,
			RefundPreference:      GRANT_PRORATED,
			SampleContentProvided: true,
		}

		data, err := json.Marshal(req)
		assert.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		assert.NoError(t, err)

		val, exists := result["consumptionPercentage"]
		assert.True(t, exists)
		assert.Equal(t, float64(50), val)
	})
}

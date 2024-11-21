package sdktime

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type instance struct {
	Time *SdkTime `json:"time"`
}

func TestSdkTime_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		Name string
		Data string
	}{
		{"format1", `{"time":"2024-05-30T14:30:00"}`},
		{"format2", `{"time":"2024-05-30T14:30:00Z"}`},
		{"format3", `{"time":"2024-05-30T14:30:00+02:00"}`},
		{"empty", `{"time":""}`},
		{"null", `{"time":null}`},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ins := &instance{}
			err := utils.Unmarshal([]byte(c.Data), ins)
			assert.Nil(t, err)
		})
	}

}

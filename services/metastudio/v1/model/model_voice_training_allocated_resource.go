package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VoiceTrainingAllocatedResource struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	ChargeMode *ChardMode `json:"charge_mode,omitempty"`

	// 资源过期时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o VoiceTrainingAllocatedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoiceTrainingAllocatedResource struct{}"
	}

	return strings.Join([]string{"VoiceTrainingAllocatedResource", string(data)}, " ")
}

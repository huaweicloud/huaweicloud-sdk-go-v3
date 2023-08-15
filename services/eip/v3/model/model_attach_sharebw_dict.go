package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachSharebwDict 共享带宽加入弹性公网IP参数
type AttachSharebwDict struct {

	// - 带宽id
	BandwidthId *string `json:"bandwidth_id,omitempty"`
}

func (o AttachSharebwDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachSharebwDict struct{}"
	}

	return strings.Join([]string{"AttachSharebwDict", string(data)}, " ")
}

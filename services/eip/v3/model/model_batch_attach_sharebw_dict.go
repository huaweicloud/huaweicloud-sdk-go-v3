package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachSharebwDict 共享带宽批量加入弹性公网IP参数
type BatchAttachSharebwDict struct {

	// - 共享带宽的id
	BandwidthId *string `json:"bandwidth_id,omitempty"`

	// - 弹性公网IP ID
	PublicipId *string `json:"publicip_id,omitempty"`
}

func (o BatchAttachSharebwDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachSharebwDict struct{}"
	}

	return strings.Join([]string{"BatchAttachSharebwDict", string(data)}, " ")
}

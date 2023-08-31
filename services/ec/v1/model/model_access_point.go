package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessPoint struct {

	// 企业连接网络接入点ID
	Id *string `json:"id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 带宽
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 绑定智能企业网关数量
	BindIegCount *int32 `json:"bind_ieg_count,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o AccessPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPoint struct{}"
	}

	return strings.Join([]string{"AccessPoint", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenPublicIp 弹性公网IP对象
type OpenPublicIp struct {

	// 弹性IP绑定类型，取值如下： auto_assign：自动绑定 not_use：暂未使用 bind_existing ：使用已有
	PublicBindType *string `json:"public_bind_type,omitempty"`

	// 弹性IP的ID
	EipId *string `json:"eip_id,omitempty"`
}

func (o OpenPublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenPublicIp struct{}"
	}

	return strings.Join([]string{"OpenPublicIp", string(data)}, " ")
}

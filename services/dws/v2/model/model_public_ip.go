package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicIp struct {

	// 弹性IP绑定类型，取值如下：  - auto_assign：自动绑定 - not_use：暂未使用 - bind_existing ：使用已有
	PublicBindType string `json:"public_bind_type" xml:"public_bind_type"`

	// 弹性IP的id
	EipId *string `json:"eip_id,omitempty" xml:"eip_id"`
}

func (o PublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIp struct{}"
	}

	return strings.Join([]string{"PublicIp", string(data)}, " ")
}

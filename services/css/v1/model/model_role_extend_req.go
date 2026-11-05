package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RoleExtendReq struct {

	// 集群扩容请求详细描述。
	Grow []RoleExtendGrowReq `json:"grow"`

	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o RoleExtendReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleExtendReq struct{}"
	}

	return strings.Join([]string{"RoleExtendReq", string(data)}, " ")
}

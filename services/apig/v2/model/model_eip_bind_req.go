package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EipBindReq struct {
	// 弹性公网IP编号

	EipId *string `json:"eip_id,omitempty"`
}

func (o EipBindReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipBindReq struct{}"
	}

	return strings.Join([]string{"EipBindReq", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpcepConnectionReq struct {
	// 期望的操作行为。

	Action string `json:"action"`
	// 终端节点ID列表（用户ID）。

	EndpointIdList []string `json:"endpointIdList"`
}

func (o UpdateVpcepConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcepConnectionReq struct{}"
	}

	return strings.Join([]string{"UpdateVpcepConnectionReq", string(data)}, " ")
}

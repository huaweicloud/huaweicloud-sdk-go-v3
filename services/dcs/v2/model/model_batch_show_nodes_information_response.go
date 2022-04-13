package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchShowNodesInformationResponse struct {
	// 查询结果的实例总数

	Count *int32 `json:"count,omitempty"`
	// 实例列表。

	Instances      *[]InstanceNodesInfoResp `json:"instances,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchShowNodesInformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowNodesInformationResponse struct{}"
	}

	return strings.Join([]string{"BatchShowNodesInformationResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodesResponse Response Object
type ShowNodesResponse struct {

	// 作业算子基本信息列表
	Body           *[]JobAndNodeInfo `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodesResponse struct{}"
	}

	return strings.Join([]string{"ShowNodesResponse", string(data)}, " ")
}

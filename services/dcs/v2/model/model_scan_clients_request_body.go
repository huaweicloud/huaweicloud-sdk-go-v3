package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScanClientsRequestBody struct {

	// 节点ID
	NodeId string `json:"node_id"`

	// 是否重新查询并保存会话列表
	CleanCache *bool `json:"clean_cache,omitempty"`
}

func (o ScanClientsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanClientsRequestBody struct{}"
	}

	return strings.Join([]string{"ScanClientsRequestBody", string(data)}, " ")
}

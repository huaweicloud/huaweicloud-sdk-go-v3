package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OfflineNodesRequestBody 隔离节点请求体。
type OfflineNodesRequestBody struct {

	// 操作类型 ,shutdown代表关机,开机是startup。
	Action string `json:"action"`

	// 节点ID列表,最多10个。
	NodeIds []string `json:"node_ids"`
}

func (o OfflineNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfflineNodesRequestBody struct{}"
	}

	return strings.Join([]string{"OfflineNodesRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchGetLabelRequestBody 获取多集群节点标签请求体
type BatchGetLabelRequestBody struct {

	// **参数解释**: 集群id列表 **取值范围**: 取值0-65535
	ClusterIds []string `json:"cluster_ids"`
}

func (o BatchGetLabelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchGetLabelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchGetLabelRequestBody", string(data)}, " ")
}

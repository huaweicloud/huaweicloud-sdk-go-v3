package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StopInstanceRequestBody struct {

	// 需要停止的节点的ID，取值不能为null，如果为空列表则停止整个实例
	NodeIds []string `json:"node_ids"`
}

func (o StopInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"StopInstanceRequestBody", string(data)}, " ")
}

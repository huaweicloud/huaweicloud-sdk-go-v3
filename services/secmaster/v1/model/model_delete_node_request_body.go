package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteNodeRequestBody struct {

	// 节点ID列表
	DeleteIds []string `json:"delete_ids"`
}

func (o DeleteNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteNodeRequestBody", string(data)}, " ")
}

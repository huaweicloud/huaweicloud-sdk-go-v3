package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteReadonlyNodeRequestBody struct {

	// 所有需要删除的节点ID。
	NodeList []string `json:"node_list"`
}

func (o DeleteReadonlyNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReadonlyNodeRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteReadonlyNodeRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteMongosNodeRequestBody struct {

	// 所有需要删除的mongos节点ID，至少保留两个mongos节点。
	NodeList []string `json:"node_list"`
}

func (o DeleteMongosNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMongosNodeRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteMongosNodeRequestBody", string(data)}, " ")
}

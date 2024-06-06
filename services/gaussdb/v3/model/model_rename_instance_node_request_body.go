package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RenameInstanceNodeRequestBody struct {

	// 节点信息列表
	NodeList []SingleNodeInfo `json:"node_list"`
}

func (o RenameInstanceNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameInstanceNodeRequestBody struct{}"
	}

	return strings.Join([]string{"RenameInstanceNodeRequestBody", string(data)}, " ")
}

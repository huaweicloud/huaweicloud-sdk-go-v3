package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEndpointRequestBody struct {

	// 终端节点名称。 取值范围：1-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）。
	Name string `json:"name"`
}

func (o UpdateEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRequestBody", string(data)}, " ")
}

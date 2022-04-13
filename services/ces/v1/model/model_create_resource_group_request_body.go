package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建资源分组，请求参数。
type CreateResourceGroupRequestBody struct {
	// 资源分组的名称；长度为1-128，只能包含0-9/a-z/A-Z/_/-或汉字；如：ResourceGroup-Test01。

	GroupName string `json:"group_name"`
	// 创建的资源分组选择一个或者多个资源。

	Resources []CreateResourceGroup `json:"resources"`
}

func (o CreateResourceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源分组修改请求体
type PutResourceGroupReq struct {

	// 资源分组名称，只能为字母、数字、汉字、-、_，最大长度为128
	GroupName string `json:"group_name"`

	// 标签动态匹配时的关联标签
	Tags *[]ResourceGroupTagRelation `json:"tags,omitempty"`
}

func (o PutResourceGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutResourceGroupReq struct{}"
	}

	return strings.Join([]string{"PutResourceGroupReq", string(data)}, " ")
}

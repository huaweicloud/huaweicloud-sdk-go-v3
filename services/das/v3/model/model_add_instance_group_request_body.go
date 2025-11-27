package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddInstanceGroupRequestBody struct {

	// 数据库类型
	DatastoreType string `json:"datastore_type"`

	// 实例组名称
	GroupName string `json:"group_name"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o AddInstanceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInstanceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"AddInstanceGroupRequestBody", string(data)}, " ")
}

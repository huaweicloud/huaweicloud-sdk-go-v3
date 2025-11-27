package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitSubnetOption 创建中转子网请求体
type CreateTransitSubnetOption struct {

	// 中转子网的名字。仅支持数字、字母、_（下划线）、-（中划线）、中文。
	Name string `json:"name"`

	// 中转子网的描述。长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	// 中转子网的子网ID。
	VirsubnetId string `json:"virsubnet_id"`

	// 中转子网的子网所属的项目ID。仅支持数字和小写字母。
	VirsubnetProjectId string `json:"virsubnet_project_id"`

	// tag标签。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateTransitSubnetOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitSubnetOption struct{}"
	}

	return strings.Join([]string{"CreateTransitSubnetOption", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJobGroup 创建会话请求参数groups的元素。
type SparkJobGroup struct {

	// 用户组名称。
	Name *string `json:"name,omitempty"`

	// 用户组资源。
	Resources *[]SparkJobResource `json:"resources,omitempty"`
}

func (o SparkJobGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobGroup struct{}"
	}

	return strings.Join([]string{"SparkJobGroup", string(data)}, " ")
}

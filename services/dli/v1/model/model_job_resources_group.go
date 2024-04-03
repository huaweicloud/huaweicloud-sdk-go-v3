package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobResourcesGroup 创建会话请求参数groups的元素。
type JobResourcesGroup struct {

	// 用户组名称。
	Name *string `json:"name,omitempty"`

	// 用户组资源。
	Resources *[]JobResource `json:"resources,omitempty"`
}

func (o JobResourcesGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResourcesGroup struct{}"
	}

	return strings.Join([]string{"JobResourcesGroup", string(data)}, " ")
}

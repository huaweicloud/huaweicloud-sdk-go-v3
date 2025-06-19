package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobResource 创建会话请求参数resources的元素。
type JobResource struct {

	// 参数解释:   资源名称 示例: group.tesddws/gsjdbc3.jar 约束限制:  无 取值范围: 无 默认取值: 无
	Name *string `json:"name,omitempty"`

	// 参数解释:   资源类型 示例: jar 约束限制:  无 取值范围: 无 默认取值: 无
	Type *string `json:"type,omitempty"`
}

func (o JobResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResource struct{}"
	}

	return strings.Join([]string{"JobResource", string(data)}, " ")
}

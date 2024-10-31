package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LabelVo 标签
type LabelVo struct {

	// uri主键
	Uri *string `json:"uri,omitempty"`

	// 逻辑region
	Region *string `json:"region,omitempty"`

	// 标签名称
	LabelName *string `json:"label_name,omitempty"`

	// 服务类型
	ServiceType *string `json:"service_type,omitempty"`

	// 所属资源类型（TestCase：用例，Task：测试套）
	ResourceType *string `json:"resource_type,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`
}

func (o LabelVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelVo struct{}"
	}

	return strings.Join([]string{"LabelVo", string(data)}, " ")
}

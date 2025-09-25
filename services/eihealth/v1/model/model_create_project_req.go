package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectReq 创建项目请求体
type CreateProjectReq struct {

	// 项目描述
	Description *string `json:"description,omitempty"`

	// 项目名称
	Name string `json:"name"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`

	// 是否核心项目
	IsCore *bool `json:"is_core,omitempty"`

	// 是否新桶, 仅气象支持该字段
	IsNewBucket *bool `json:"is_new_bucket,omitempty"`

	// 桶名, 仅气象支持该字段
	BucketName *string `json:"bucket_name,omitempty"`

	// **参数解释**： 空间桶来源租户，TERMINATE表示最终租户桶。 **约束限制**： 不涉及 **取值范围**： * TERMINATE： 最终租户桶。 * RESOURCE： 资源租户桶。 **默认取值**： TERMINATE
	StorageTenant *string `json:"storage_tenant,omitempty"`
}

func (o CreateProjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectReq struct{}"
	}

	return strings.Join([]string{"CreateProjectReq", string(data)}, " ")
}

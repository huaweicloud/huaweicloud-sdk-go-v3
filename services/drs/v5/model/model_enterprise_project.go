package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseProject 企业项目信息。
type EnterpriseProject struct {

	// 企业项目ID，0表示默认企业项目。
	Id *string `json:"id,omitempty"`

	// 企业项目名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 企业项目状态。 - 1：启用 - 2：停用
	Status *string `json:"status,omitempty"`

	// 创建时间，例如：2023-01-20T07:18:26Z
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间，例如：2023-03-01T09:42:02Z
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o EnterpriseProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProject struct{}"
	}

	return strings.Join([]string{"EnterpriseProject", string(data)}, " ")
}

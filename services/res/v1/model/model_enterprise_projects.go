package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseProjects struct {

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 企业项目id。
	Id *string `json:"id,omitempty"`

	// 企业项目名称。
	Name *string `json:"name,omitempty"`

	// 状态。
	Status *int32 `json:"status,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o EnterpriseProjects) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjects struct{}"
	}

	return strings.Join([]string{"EnterpriseProjects", string(data)}, " ")
}

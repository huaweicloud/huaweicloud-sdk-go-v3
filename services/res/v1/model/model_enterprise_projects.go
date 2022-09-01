package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseProjects struct {

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 企业项目id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 企业项目名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 状态。
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`
}

func (o EnterpriseProjects) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjects struct{}"
	}

	return strings.Join([]string{"EnterpriseProjects", string(data)}, " ")
}

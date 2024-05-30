package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineCreateReqEnterpriseProject 企业项目信息
type EngineCreateReqEnterpriseProject struct {

	// 企业项目id
	Id *string `json:"id,omitempty"`

	// 企业项目名称
	Name *string `json:"name,omitempty"`

	// 企业项目描述
	Description *string `json:"description,omitempty"`

	// 企业项目状态
	Status *int32 `json:"status,omitempty"`

	// 企业项目创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 企业项目升级时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 企业项目标签
	Label *string `json:"label,omitempty"`
}

func (o EngineCreateReqEnterpriseProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineCreateReqEnterpriseProject struct{}"
	}

	return strings.Join([]string{"EngineCreateReqEnterpriseProject", string(data)}, " ")
}

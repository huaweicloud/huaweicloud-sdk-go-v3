package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLakeFormationInstanceResponse Response Object
type UpdateLakeFormationInstanceResponse struct {

	// LakeFormation实例Id
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 企业项目Id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 逻辑多租和物理多租的判断
	Shared *bool `json:"shared,omitempty"`

	// 实例创建时间戳
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 实例更新时间戳
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 实例状态
	Status *string `json:"status,omitempty"`

	// 是否在回收站
	InRecycleBin   *bool `json:"in_recycle_bin,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateLakeFormationInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLakeFormationInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateLakeFormationInstanceResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VariableRes struct {
	ByOrder *int32 `json:"by_order,omitempty"`

	Category *string `json:"category,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	CreateTimeStamp *int64 `json:"create_time_stamp,omitempty"`

	CreateTimeString *string `json:"create_time_string,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	CurrentPermission *string `json:"currentPermission,omitempty"`

	Description *string `json:"description,omitempty"`

	DynamicParamFlag *bool `json:"dynamicParamFlag,omitempty"`

	FunctionParams *string `json:"functionParams,omitempty"`

	GroupId *string `json:"groupId,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	IsSensitiveInfo *bool `json:"isSensitiveInfo,omitempty"`

	IsSensitiveModified *bool `json:"isSensitiveModified,omitempty"`

	Locked *int32 `json:"locked,omitempty"`

	Name *string `json:"name,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	NodeType *int32 `json:"node_type,omitempty"`

	ParentId *string `json:"parent_id,omitempty"`

	ParentNodeId *string `json:"parent_node_id,omitempty"`

	Property *string `json:"property,omitempty"`

	Region *string `json:"region,omitempty"`

	SensitiveInfoSetterTime *string `json:"sensitiveInfoSetterTime,omitempty"`

	SensitiveInfoSetterUser *string `json:"sensitiveInfoSetterUser,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`

	Type *string `json:"type,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	UpdateTimeStamp *int64 `json:"update_time_stamp,omitempty"`

	UpdateTimeString *string `json:"update_time_string,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`

	VariableType *int32 `json:"variableType,omitempty"`
}

func (o VariableRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VariableRes struct{}"
	}

	return strings.Join([]string{"VariableRes", string(data)}, " ")
}

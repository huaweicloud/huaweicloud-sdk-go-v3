package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointBriefInfo Endpoint简要信息
type EndpointBriefInfo struct {

	// 可见性：  - PRIVATE：私有  - PUBLIC：公共  默认为PRIVATE
	Visibility *string `json:"visibility,omitempty"`

	// Endpoint Id，32~36位的英文、数字、短横组合。
	Id string `json:"id"`

	// 一个Endpoint名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	Type *EndpointType `json:"type"`

	Status *EndpointStatus `json:"status"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time"`

	Owner *User `json:"owner"`

	Cap *CapRef `json:"cap,omitempty"`

	ReservedResource *ReservedResource `json:"reserved_resource,omitempty"`

	RayResource *RayResourceInfo `json:"ray_resource,omitempty"`

	// 失败状态时的错误编码
	ErrorCode *string `json:"error_code,omitempty"`

	// 失败状态时的错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o EndpointBriefInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointBriefInfo struct{}"
	}

	return strings.Join([]string{"EndpointBriefInfo", string(data)}, " ")
}

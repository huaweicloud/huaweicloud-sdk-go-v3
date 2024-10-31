package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServiceInstanceDetailResponse Response Object
type ShowServiceInstanceDetailResponse struct {

	// 可见性：  - PRIVATE：私有  - PUBLIC：公共  默认为PRIVATE
	Visibility *string `json:"visibility,omitempty"`

	Source *SourceRef `json:"source,omitempty"`

	// Service Instance的ID
	Id *string `json:"id,omitempty"`

	// 一个Service Instance的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// endpoint空间id
	EndpointId *string `json:"endpoint_id,omitempty"`

	Status *StatusEnum `json:"status,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 运行时长
	Duration *int64 `json:"duration,omitempty"`

	CreateUser *User `json:"create_user,omitempty"`

	Type *ServiceType `json:"type,omitempty"`

	UpdateUser *User `json:"update_user,omitempty"`

	Url *Url `json:"url,omitempty"`

	Config *AppInstanceConfig `json:"config,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowServiceInstanceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServiceInstanceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowServiceInstanceDetailResponse", string(data)}, " ")
}

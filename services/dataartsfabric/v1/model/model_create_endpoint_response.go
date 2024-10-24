package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEndpointResponse Response Object
type CreateEndpointResponse struct {

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

	// 调用地址
	Urls *[]Url `json:"urls,omitempty"`

	LogConfig *LogConfigInfo `json:"log_config,omitempty"`

	// 引擎实例id列表
	BusinessEngineInstanceIds *[]string `json:"business_engine_instance_ids,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateEndpointResponse", string(data)}, " ")
}

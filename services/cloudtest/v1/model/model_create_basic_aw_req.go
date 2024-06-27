package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateBasicAwReq struct {

	// body请求体的类型：text(包含JSON，参数内部区分)、form
	BodyParamType *string `json:"body_param_type,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 默认检查点List
	DftCheckPointList *[]CheckPoint `json:"dft_check_point_list,omitempty"`

	// 默认请求头参数对象
	DftCustomHeader *[]AwParam `json:"dft_custom_header,omitempty"`

	// 定义的变量信息
	DftVariableList *[]AwVariable `json:"dft_variable_list,omitempty"`

	ExtraInfo *ExtraInfo `json:"extra_info,omitempty"`

	// 组名
	GroupName *string `json:"group_name,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 参数类型和参数默认值对应List
	ParamTypeAndDftValue *[]AwParam `json:"param_type_and_dft_value,omitempty"`
}

func (o CreateBasicAwReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBasicAwReq struct{}"
	}

	return strings.Join([]string{"CreateBasicAwReq", string(data)}, " ")
}

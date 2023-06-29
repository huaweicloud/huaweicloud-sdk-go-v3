package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiMock mock后端详情
type ApiMock struct {

	// 描述信息。长度不超过255个字符 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 返回结果
	ResultContent *string `json:"result_content,omitempty"`

	// 版本。字符长度不超过64
	Version *string `json:"version,omitempty"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// 后端状态   - 1： 有效
	Status *int32 `json:"status,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o ApiMock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiMock struct{}"
	}

	return strings.Join([]string{"ApiMock", string(data)}, " ")
}

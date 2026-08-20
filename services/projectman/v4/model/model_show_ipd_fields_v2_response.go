package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdFieldsV2Response Response Object
type ShowIpdFieldsV2Response struct {

	// 响应状态码。标识查询工作项实例字段列表请求的处理结果。
	Status *string `json:"status,omitempty"`

	// 响应消息。请求失败时包含详细错误信息，可用于问题排查。
	Message *string `json:"message,omitempty"`

	// 字段列表结果。返回创建工作项实例时可选用的字段配置信息，包含字段ID、编码、名称、类型等属性，包含系统字段和项目自定义字段。
	Result         *[]FieldVo `json:"result,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowIpdFieldsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdFieldsV2Response struct{}"
	}

	return strings.Join([]string{"ShowIpdFieldsV2Response", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetadataResponse 响应数据。
type MetadataResponse struct {

	// 组件ID。
	Id *string `json:"id,omitempty"`

	// 组件名称。
	Name *string `json:"name,omitempty"`

	// 组件附加属性。 - log_group_id：LTS日志组的ID。 - log_stream_id：LTS日志流的ID。 - version：组件版本。
	Annotations map[string]string `json:"annotations,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o MetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataResponse struct{}"
	}

	return strings.Join([]string{"MetadataResponse", string(data)}, " ")
}

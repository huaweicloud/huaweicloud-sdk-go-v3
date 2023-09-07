package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomConnectorFromOpenapiResponse Response Object
type CreateCustomConnectorFromOpenapiResponse struct {

	// 自定义连接器版本ID
	ConnectorVersionId *string `json:"connector_version_id,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	Description *string `json:"description,omitempty"`

	// 自定义连接器ID
	Id *string `json:"id,omitempty"`

	// 自定义连接器名称
	Name *string `json:"name,omitempty"`

	// 修改时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCustomConnectorFromOpenapiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomConnectorFromOpenapiResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomConnectorFromOpenapiResponse", string(data)}, " ")
}

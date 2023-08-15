package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeAppResponse Response Object
type ShowEdgeAppResponse struct {

	// 应用id
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 最新发布版本
	LastPublishedVersion *string `json:"last_published_version,omitempty"`

	// 应用类型SYSTEM_REQUIRED|SYSTEM_OPTIONAL|USER
	AppType *string `json:"app_type,omitempty"`

	// 应用类型DATA_PROCESSING|PROTOCOL_PARSING
	FunctionType *string `json:"function_type,omitempty"`

	// 部署类型docker|process
	DeployType *string `json:"deploy_type,omitempty"`

	// 驱动协议类型OPCUA|Modbus-TCP
	Protocol *string `json:"protocol,omitempty"`

	// 应用名称
	EdgeAppName    *string `json:"edge_app_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEdgeAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeAppResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeAppResponse", string(data)}, " ")
}

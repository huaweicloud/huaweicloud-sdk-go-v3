package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupResponse Response Object
type ShowGroupResponse struct {

	// 是否可以消费。
	Enabled *bool `json:"enabled,omitempty"`

	// 是否广播。
	Broadcast *bool `json:"broadcast,omitempty"`

	// 关联的代理列表。
	Brokers *[]string `json:"brokers,omitempty"`

	// 消费组名称。
	Name *string `json:"name,omitempty"`

	// 消费组描述。
	GroupDesc *string `json:"group_desc,omitempty"`

	// 最大重试次数。
	RetryMaxTime *int32 `json:"retry_max_time,omitempty"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 权限。
	Permissions    *[]string `json:"permissions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupResponse", string(data)}, " ")
}

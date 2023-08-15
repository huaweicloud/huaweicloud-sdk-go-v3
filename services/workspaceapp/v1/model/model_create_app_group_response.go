package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppGroupResponse Response Object
type CreateAppGroupResponse struct {

	// 应用组ID
	Id *string `json:"id,omitempty"`

	// 应用组名称
	Name *string `json:"name,omitempty"`

	// 应用服务器组ID
	AppServerGroupId *string `json:"app_server_group_id,omitempty"`

	// 应用服务器组名称
	AppServerGroupName *string `json:"app_server_group_name,omitempty"`

	// 应用组描述
	Description *string `json:"description,omitempty"`

	AuthorizationType *AuthorizationTypeEnum `json:"authorization_type,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 发布时间
	CreateAt *sdktime.SdkTime `json:"create_at,omitempty"`

	// 应用数量
	AppCount       *int32 `json:"app_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateAppGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateAppGroupResponse", string(data)}, " ")
}

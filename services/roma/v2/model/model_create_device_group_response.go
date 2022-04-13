package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDeviceGroupResponse struct {
	// 权限

	Permissions *[]string `json:"permissions,omitempty"`
	// 父分组ID

	ParentId *int32 `json:"parent_id,omitempty"`
	// 分组ID

	Id *int32 `json:"id,omitempty"`
	// 分组名称，支持中文，英文大小写，数字，下划线和中划线,长度2-64

	Name *string `json:"name,omitempty"`
	// 分组描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 分组归属应用ID

	AppId *string `json:"app_id,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`
	// 创建时间

	CreatedDatetime *string `json:"created_datetime,omitempty"`
	// 最后修改时间

	LastUpdatedDatetime *string `json:"last_updated_datetime,omitempty"`
	// 应用名称

	AppName        *string `json:"app_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDeviceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateDeviceGroupResponse", string(data)}, " ")
}

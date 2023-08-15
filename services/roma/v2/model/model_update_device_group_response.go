package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeviceGroupResponse Response Object
type UpdateDeviceGroupResponse struct {

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
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateDeviceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceGroupResponse", string(data)}, " ")
}

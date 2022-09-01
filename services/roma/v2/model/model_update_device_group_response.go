package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDeviceGroupResponse struct {

	// 权限
	Permissions *[]string `json:"permissions,omitempty" xml:"permissions"`

	// 父分组ID
	ParentId *int32 `json:"parent_id,omitempty" xml:"parent_id"`

	// 分组ID
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 分组名称，支持中文，英文大小写，数字，下划线和中划线,长度2-64
	Name *string `json:"name,omitempty" xml:"name"`

	// 分组描述，长度0-200
	Description *string `json:"description,omitempty" xml:"description"`

	// 分组归属应用ID
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	CreatedUser *CreatedUser `json:"created_user,omitempty" xml:"created_user"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty" xml:"last_updated_user"`

	// 创建时间
	CreatedDatetime *string `json:"created_datetime,omitempty" xml:"created_datetime"`

	// 最后修改时间
	LastUpdatedDatetime *string `json:"last_updated_datetime,omitempty" xml:"last_updated_datetime"`

	// 应用名称
	AppName        *string `json:"app_name,omitempty" xml:"app_name"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDeviceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceGroupResponse", string(data)}, " ")
}

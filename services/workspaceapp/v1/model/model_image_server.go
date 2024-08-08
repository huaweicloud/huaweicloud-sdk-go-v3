package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageServer struct {

	// 实例的唯一标识。
	Id *string `json:"id,omitempty"`

	// 镜像实例名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	ImageRef *ImageRef `json:"image_ref,omitempty"`

	// APS服务器组ID。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 应用组ID。
	AppGroupId *string `json:"app_group_id,omitempty"`

	// APS实例ID。
	ServerId *string `json:"server_id,omitempty"`

	// ECS服务器ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 镜像产物唯一标识。
	ImageId *string `json:"image_id,omitempty"`

	Status *ImageServerStatus `json:"status,omitempty"`

	// 应用组授权用户， * 限制用户类型：'USER' - 用户
	AuthorizeAccounts *[]ImageAccountInfo `json:"authorize_accounts,omitempty"`

	// 镜像实例创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 企业项目ID,仅企业项目需配置(字段为空或者0表示使用默认default企业项目)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ImageServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageServer struct{}"
	}

	return strings.Join([]string{"ImageServer", string(data)}, " ")
}

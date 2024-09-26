package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetwork 中心网络。
type CentralNetwork struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	State *CentralNetworkStateEnum `json:"state"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例标签。
	Tags *[]Tag `json:"tags,omitempty"`

	// 中心网络默认平面的ID。
	DefaultPlaneId string `json:"default_plane_id"`

	// 中心网平面列表。
	Planes *[]CentralNetworkPlane `json:"planes,omitempty"`

	// 中心网ER实例列表。
	ErInstances *[]CentralNetworkErInstance `json:"er_instances,omitempty"`

	// 中心网ER连接列表。
	Connections *[]CentralNetworkConnectionInfo `json:"connections,omitempty"`
}

func (o CentralNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetwork struct{}"
	}

	return strings.Join([]string{"CentralNetwork", string(data)}, " ")
}

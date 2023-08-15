package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransitIp 中转子网IP的响应体。
type TransitIp struct {

	// 中转IP的ID。
	Id string `json:"id"`

	// 项目的ID。
	ProjectId string `json:"project_id"`

	// 中转IP的网络接口ID。
	NetworkInterfaceId string `json:"network_interface_id"`

	// 中转IP的地址。
	IpAddress string `json:"ip_address"`

	// 中转IP的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 中转IP的更新时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 当前租户子网的ID。
	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	// 标签列表。
	Tags *[]PrivateTag `json:"tags,omitempty"`

	// 中转IP绑定的私网NAT网关实例的ID。
	GatewayId string `json:"gateway_id"`

	// 企业项目ID。创建中转IP时，关联的企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o TransitIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransitIp struct{}"
	}

	return strings.Join([]string{"TransitIp", string(data)}, " ")
}

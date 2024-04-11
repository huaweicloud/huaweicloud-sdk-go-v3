package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudPhoneServerDetailResponse Response Object
type ShowCloudPhoneServerDetailResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 云手机服务器名称， 不超过65字符，只支持英文字母、数字、汉字、下划线和中划线。
	ServerName *string `json:"server_name,omitempty"`

	// 云手机服务器所在的可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 云手机服务器的唯一标识，不超过32个字节。
	ServerId *string `json:"server_id,omitempty"`

	// 云手机服务器规格名称，不超过64个字节。
	ServerModelName *string `json:"server_model_name,omitempty"`

	// 云手机规格名称，不超过64个字节。
	PhoneModelName *string `json:"phone_model_name,omitempty"`

	// 连接云手机所使用的密钥对的名称，不超过64个字节。
	KeypairName *string `json:"keypair_name,omitempty"`

	// 服务器状态。 - 0、1、3、4：创建中 - 2：异常 - 5：正常 - 8：冻结 - 10：关机 - 11：关机中 - 12：关机失败 - 13：开机中
	Status *int32 `json:"status,omitempty"`

	// 云手机服务器所属虚拟私有云（简称VPC）的ID。网络版本 network_version 取值为“v1”时，表示云手机服务器所属资源租户的VPC ID；取值为“v2”时，表示租户创建服务器时指定 VPC 的 VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 云手机服务器所属虚拟私有云（简称VPC）的网段。网络版本 network_version 取值为“v1”时，表示云手机服务器所属资源租户的VPC CIDR；取值为“v2”时，表示租户创建服务器时指定 VPC 的 VPC CIDR。
	Cidr *string `json:"cidr,omitempty"`

	// 云手机服务器所属虚拟私有云（简称VPC）的网段。网络版本 network_version 取值为“v1”时，表示云手机服务器所属资源租户的VPC CIDR；取值为“v2”时，表示租户创建服务器时指定 VPC 的 VPC CIDR。
	VpcCidr *string `json:"vpc_cidr,omitempty"`

	// 云手机服务器所属子网的ID。仅在网络版本 network_version 取值为“v2”时，该取值表示租户创建服务器时指定子网的 ID；网络版本取值为“v1”时，该字段表示云手机服务器所属资源租户的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 云手机服务器所属子网的网段。网络版本 network_version取值为“v2”时，表示租户创建服务器时指定子网的 CIDR；取值为“v1”时，表示云手机服务器所属资源租户的子网CIDR。
	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	// 云手机服务器的项目ID。
	ResourceProjectId *string `json:"resource_project_id,omitempty"`

	Metadata *ShowCloudPhoneServerDetailResponseBodyMetadata `json:"metadata,omitempty"`

	// 云手机服务器的IP相关信息。
	Addresses *[]Address `json:"addresses,omitempty"`

	// 云手机服务器带宽信息的结构体数组。
	BandWidths *[]Bandwidth `json:"band_widths,omitempty"`

	// 云手机服务器卷信息的结构体数组。
	Volumes *[]Volume `json:"volumes,omitempty"`

	ShareVolumeInfo *ShowCloudPhoneServerDetailResponseBodyShareVolumeInfo `json:"share_volume_info,omitempty"`

	// 是否为自定义网络的云手机服务器标识。\"v1\"，非自定义网络的云手机服务器。\"v2\"，自定义网络的云手机服务器。支持按照网络版本字段进行筛选。
	NetworkVersion *string `json:"network_version,omitempty"`

	// 云手机服务器所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 服务器扩展网卡绑定的安全组信息。 系统定义网络的服务器，该字段返回为空列表。
	SecurityGroups *[]string `json:"security_groups,omitempty"`

	// 创建时间， 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间， 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ。
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCloudPhoneServerDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudPhoneServerDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCloudPhoneServerDetailResponse", string(data)}, " ")
}

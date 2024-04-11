package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeCloudPhoneServerRequestBody struct {

	// 云手机规格，不超过64个字节。
	PhoneModelName string `json:"phone_model_name"`

	// 云手机镜像ID，不超过32个字节。
	ImageId string `json:"image_id"`

	// 密钥对名称，不超过64个字节，用于云手机ADB登录。
	KeypairName *string `json:"keypair_name,omitempty"`

	// 云手机启用的应用端口，云手机服务会做端口转发。
	Ports *[]Port `json:"ports,omitempty"`

	ExtendParam *ChangeCloudPhoneServerRequestBodyExtendParam `json:"extend_param,omitempty"`

	// 租户自定义的VPC ID，为待创建的云服务器所属的虚拟私有云（简称VPC），需要指定已创建VPC的ID，UUID格式。
	TenantVpcId string `json:"tenant_vpc_id"`

	// 租户自定义的网卡的结构体，为待创建的云服务器的网卡信息。
	Nics []Nic `json:"nics"`

	PublicIp *ChangeCloudPhoneServerRequestBodyPublicIp `json:"public_ip"`

	// 多少个手机共用一个vip。默认为手机开数，表示所有手机共享1个vip。取值范围：1到手机规格开数。
	PhoneCountPerIp *int32 `json:"phone_count_per_ip,omitempty"`

	PhoneDataVolume *CreateNet2CloudPhoneServerRequestBodyPhoneDataVolume `json:"phone_data_volume,omitempty"`

	ServerShareDataVolume *ChangeCloudPhoneServerRequestBodyServerShareDataVolume `json:"server_share_data_volume,omitempty"`

	BandWidth *ChangeCloudPhoneServerRequestBodyBandWidth `json:"band_width"`
}

func (o ChangeCloudPhoneServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerRequestBody", string(data)}, " ")
}

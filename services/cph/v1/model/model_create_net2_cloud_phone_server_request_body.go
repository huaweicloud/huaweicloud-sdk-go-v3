package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNet2CloudPhoneServerRequestBody struct {

	// 云手机服务器名称  不超过60个字符，只支持英文字母、数字、汉字、下划线和中划线。  批量购买会在服务器名称后自动添加序号，比如设置此参数为server-1，那么创建的云手机服务器名称会自动按序增加数字后缀，比如为server-1-0001
	ServerName string `json:"server_name"`

	// 云手机服务器规格，不超过64个字节
	ServerModelName string `json:"server_model_name"`

	// 云手机规格，不超过64个字节
	PhoneModelName string `json:"phone_model_name"`

	// 云手机镜像ID，不超过32个字节
	ImageId string `json:"image_id"`

	// 购买的云手机服务器个数，最多可购买10台
	Count int32 `json:"count"`

	// 密钥对名称，不超过64个字节，用于云手机ADB登录
	KeypairName *string `json:"keypair_name,omitempty"`

	// 云手机启用的应用端口，云手机服务会做端口转发
	Ports *[]Port `json:"ports,omitempty"`

	ExtendParam *CreateNet2CloudPhoneServerRequestBodyExtendParam `json:"extend_param"`

	// 租户自定义的VPC ID，为待创建的云服务器所属的虚拟私有云（简称VPC），需要指定已创建VPC的ID，UUID格式
	TenantVpcId string `json:"tenant_vpc_id"`

	// 租户自定义的网卡的结构体，为待创建的云服务器的网卡信息
	Nics []Nic `json:"nics"`

	PublicIp *CreateNet2CloudPhoneServerRequestBodyPublicIp `json:"public_ip"`

	BandWidth *CreateNet2CloudPhoneServerRequestBodyBandWidth `json:"band_width"`

	// 待创建云服务器所在的可用区，需要指定可用区（AZ）的名称
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o CreateNet2CloudPhoneServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBody", string(data)}, " ")
}

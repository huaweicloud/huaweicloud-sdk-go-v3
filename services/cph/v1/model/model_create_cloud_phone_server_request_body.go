package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCloudPhoneServerRequestBody struct {

	// 云手机服务器名称 不超过60个字符，只支持英文字母、数字、汉字、下划线和中划线。 批量购买会在服务器名称后自动添加序号，比如设置此参数为server-1，那么创建的云手机服务器名称会自动按序增加数字后缀，比如为server-1-0001
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
	Ports *[]CreateCloudPhoneServerRequestBodyPorts `json:"ports,omitempty"`

	BandWidth *CreateCloudPhoneServerRequestBodyBandWidth `json:"band_width"`

	ExtendParam *CreateCloudPhoneServerRequestBodyExtendParam `json:"extend_param"`

	// (已废弃)是否开启VNC方式登录云手机。 - 为\"true\"时开启（忽略大小写）。 - 为其他，则不开启
	VncEnable *string `json:"vnc_enable,omitempty"`

	// 服务器的子网信息，第一次购买系统会自动创建172.31.0.0/16的子网。需要自定义子网的客户，需要全部通过API购买，设置的子网，必须是子网的格式且和已有子网不能重叠
	SubnetCidr *string `json:"subnet_cidr,omitempty"`
}

func (o CreateCloudPhoneServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneServerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneServerRequestBody", string(data)}, " ")
}

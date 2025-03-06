package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudPhoneSingleServerRequestBody 创建云手机裸服务器请求体
type CreateCloudPhoneSingleServerRequestBody struct {

	// 云手机裸服务器名称，不超过60个字符，只支持英文字母、数字、汉字、下划线和中划线。批量购买会在云手机裸服务器名称后自动添加序号，比如设置此参数为server-1，那么创建的云手机裸服务器名称会自动按序增加数字后缀，比如为server-1-0001。
	ServerName string `json:"server_name"`

	// 云手机裸服务器规格名称。
	ServerModelName string `json:"server_model_name"`

	// 购买的云手机裸服务器个数，最多可购买10台。
	Count int32 `json:"count"`

	OrderParam *CreateCloudPhoneSingleServerRequestBodyOrderParam `json:"order_param"`

	// 租户自定义的网卡的结构体，为待创建的云手机裸服务器的网卡信息。
	Nics []NicForSingleServer `json:"nics"`

	PublicIp *CreateCloudPhoneSingleServerRequestBodyPublicIp `json:"public_ip"`

	// 待创建云手机裸服务器所在的可用区（AZ）的名称。如上海一可用区1为cn-east-3a。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	DataVolume *CreateCloudPhoneSingleServerRequestBodyDataVolume `json:"data_volume,omitempty"`

	// 指定登录云手机裸服务器已有密钥的名称。
	KeypairName string `json:"keypair_name"`

	// 企业项目ID。 该字段不传（或传为字符串“0”），则将资源绑定给默认企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 云手机裸服务器镜像ID，不超过36个字节。未指定时使用默认云手机裸服务器镜像。
	ImageId *string `json:"image_id,omitempty"`
}

func (o CreateCloudPhoneSingleServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneSingleServerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneSingleServerRequestBody", string(data)}, " ")
}

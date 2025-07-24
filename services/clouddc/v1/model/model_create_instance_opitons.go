package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceOpitons 创建实例请求参数
type CreateInstanceOpitons struct {

	// 设置实例主机名称。取值范围：只能由中文字符、英文字母（a~z，A~Z）、数字（0~9）、下划线（_）、中划线（-）、点（.）组成，且长度为[1-63]个字符。
	Name string `json:"name"`

	// 镜像ID
	ImageId string `json:"image_id"`

	// 创建网卡所属的 VPC ID，可通过 VPC API 查询：https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html。
	VpcId string `json:"vpc_id"`

	// 指定裸机实例的网卡信息。  约束： 一个裸机实例最多挂载2个网卡，参数中第一个网卡会作为裸机实例的主网卡。若用户指定了多组网卡参数，需保证各组参数都属于同一VPC。
	NetworkInterfaces []NetworkInterface `json:"network_interfaces"`

	// 设置实例的管理员账户初始登录密码，其中，Linux管理员账户为root，Windows管理员账户为Administrator。
	Password string `json:"password"`

	Metadata *[]map[string]string `json:"metadata,omitempty"`

	// 裸机实例的描述信息，默认为空字符串。
	Description *string `json:"description,omitempty"`

	Placement *CreateInstanceOpitonsPlacement `json:"placement,omitempty"`
}

func (o CreateInstanceOpitons) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceOpitons struct{}"
	}

	return strings.Join([]string{"CreateInstanceOpitons", string(data)}, " ")
}

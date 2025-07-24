package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Instance 裸机实例
type Instance struct {

	// UUID（Universally Unique Identifier）是一个 128 位的数字，通常以 32 个十六进制数字的形式表示，分为 5 组，用连字符分隔。具体格式如下：  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx 其中：  每个 x 是一个十六进制数字（0-9 或 a-f）。 5 组的长度分别是：8 位、4 位、4 位、4 位 和 12 位。 为了匹配这种格式的 UUID，可以使用以下正则表达式：  regex ^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$
	Id string `json:"id"`

	// instance name
	Name string `json:"name"`

	// VPC ID
	VpcId string `json:"vpc_id"`

	// 指定裸机实例的网卡信息。  约束：  一个裸机实例最多挂载2个网卡，参数中第一个网卡会作为裸机实例的主网卡。若用户指定了多组网卡参数，需保证各组参数都属于同一VPC。
	NetworkInterfaces *[]NetworkInterface `json:"network_interfaces,omitempty"`

	// 标签
	Tags []Tag `json:"tags"`

	Image *Image `json:"image"`

	// 云服务器描述信息，默认为空字符串。
	Description *string `json:"description,omitempty"`
}

func (o Instance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instance struct{}"
	}

	return strings.Join([]string{"Instance", string(data)}, " ")
}

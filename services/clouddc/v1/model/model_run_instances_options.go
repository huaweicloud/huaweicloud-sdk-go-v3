package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunInstancesOptions 创建实例请求参数
type RunInstancesOptions struct {

	// 设置实例主机名称。取值范围：只能由中文字符、英文字母（a~z，A~Z）、数字（0~9）、下划线（_）、中划线（-）、点（.）组成，且长度为[1-63]个字符。
	Name string `json:"name"`

	// 镜像ID
	ImageId string `json:"image_id"`

	// 创建网卡所属的 VPC ID，可通过 VPC API 查询：https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html。
	VpcId string `json:"vpc_id"`

	// 指定裸机实例的网卡信息。  约束： 一个裸机实例最多挂载2个网卡，参数中第一个网卡会作为裸机实例的主网卡。若用户指定了多组网卡参数，需保证各组参数都属于同一VPC。
	NetworkInterfaces []RunInstancesOptionsNetworkInterfaces `json:"network_interfaces"`

	// 设置实例的管理员账户初始登录密码，其中，Linux管理员账户为root，Windows管理员账户为Administrator。
	Password string `json:"password"`

	// 创建裸机实例的元数据。  可以通过元数据自定义键值对。   说明： 如果元数据中包含了敏感数据，您应当采取适当的措施来保护敏感数据，比如限制访问范围、加密等。 最多可注入10对键值（Key/Value）。 主键（Key）只能由大写字母（A-Z）、小写字母（a-z）、数字（0-9）、中划线（-）、下划线（_）、冒号（:）、空格（ ）和小数点（.）组成，长度为[1-255]个字符。     值（value）最大长度为255个字符。
	Metadata map[string]string `json:"metadata,omitempty"`

	// 裸机实例的描述信息，默认为空字符串。
	Description *string `json:"description,omitempty"`

	Placement *RunInstancesOptionsPlacement `json:"placement,omitempty"`

	// 必须成功启动的最小实例数量。如果无法满足此数量，整个请求失败（不会启动任何实例）。
	MinCount *int32 `json:"min_count,omitempty"`

	// 允许启动的最大实例数量。尝试启动最多该数量的实例，但实际启动数可能介于 min_count 和 max_count 之间
	MaxCount *int32 `json:"max_count,omitempty"`
}

func (o RunInstancesOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunInstancesOptions struct{}"
	}

	return strings.Join([]string{"RunInstancesOptions", string(data)}, " ")
}

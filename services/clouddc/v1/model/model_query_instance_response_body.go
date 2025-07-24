package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryInstanceResponseBody 查询实例应答
type QueryInstanceResponseBody struct {

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

	State *InstanceState `json:"state"`

	// 创建裸机实例的元数据。  可以通过元数据自定义键值对。   说明： 如果元数据中包含了敏感数据，您应当采取适当的措施来保护敏感数据，比如限制访问范围、加密等。 最多可注入10对键值（Key/Value）。 主键（Key）只能由大写字母（A-Z）、小写字母（a-z）、数字（0-9）、中划线（-）、下划线（_）、冒号（:）、空格（ ）和小数点（.）组成，长度为[1-255]个字符。     值（value）最大长度为255个字符。
	Metadata map[string]string `json:"metadata,omitempty"`

	// 创建实例过程中待注入实例自定义数据。支持注入文本、文本文件。   说明： user_data的值为base64编码之后的内容。 注入内容（编码之前的内容）最大长度为32K。
	UserData *string `json:"user_data,omitempty"`

	// UUID（Universally Unique Identifier）是一个 128 位的数字，通常以 32 个十六进制数字的形式表示，分为 5 组，用连字符分隔。具体格式如下：  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx 其中：  每个 x 是一个十六进制数字（0-9 或 a-f）。 5 组的长度分别是：8 位、4 位、4 位、4 位 和 12 位。 为了匹配这种格式的 UUID，可以使用以下正则表达式：  regex ^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$
	ServerId string `json:"server_id"`

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`

	// 启动时间
	LaunchedAt string `json:"launched_at"`

	Error *ErrorStatus `json:"error,omitempty"`
}

func (o QueryInstanceResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryInstanceResponseBody struct{}"
	}

	return strings.Join([]string{"QueryInstanceResponseBody", string(data)}, " ")
}

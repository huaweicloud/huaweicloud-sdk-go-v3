package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Condition cc规则防护规则限速条件
type Condition struct {

	// 字段类型 url：路径 ip：IPv4 ipv6：IPv6 params：Params cookie：Cookie header：Header response_code：Response Code
	Category *string `json:"category,omitempty"`

	// 条件列表逻辑匹配内容。 当匹配逻辑为exist或not_exist时，contents必须为空，其他情况下contents必填且长度不超过2048 当category为response_code时，contents状态码为200~599，正则为 ^(?:[2-5]\\d{2})$ 当匹配逻辑包含\"len\"时，contents必须为0~65535的整数；当匹配逻辑包含\"num\"时，contents必须为0~512的整数
	Contents *[]string `json:"contents,omitempty"`

	// 子字段 当字段类型为ip或ipv6时，index必填且必须为：client-ip：客户端IP、x-forwarded-for：X-Forwarded-For、TCP连接IP: $remote_addr 当字段类型（category）选择“params”、“cookie”、“header”时，请根据实际需求配置子字段且该参数必填。 当匹配逻辑为num_greater、num_less、num_equal、num_not_equal时，子字段必须为空 当子字段不为空时，最大长度不超过2048
	Index *string `json:"index,omitempty"`

	// 条件列表匹配逻辑。 如果字段类型category是url，匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 len_greater、 len_less、len_equal或者len_not_equal 如果字段类型category是ip、ipv6或response_code，匹配逻辑可以为： equal、not_equal 如果字段类型category是params、cookie或者header, 匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 len_greater、 len_less、len_equal、len_not_equal、num_greater、num_less、num_equal、num_not_equal、exist或者not_exist
	LogicOperation *string `json:"logic_operation,omitempty"`
}

func (o Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}

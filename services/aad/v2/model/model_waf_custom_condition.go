package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WafCustomCondition cc规则防护规则限速条件
type WafCustomCondition struct {

	// 字段类型 枚举值： url：路径 ip：IPv4 user-agent: User Agent method: Method referer: Referer params：Params cookie：Cookie header：Header request_line: Request Line request: Request
	Category string `json:"category"`

	// 子字段 - 当字段类型为url，user-agent、refer、request_line、method、request时，不需要传index参数 - 当字段类型为ip或ipv6时，index必填且必须为以下值：    client-ip：客户端IP    x-forwarded-for：X-Forwarded-For    TCP连接IP: $remote_addr - 当字段类型（category）选择“params”、“cookie”、“header”时，请根据实际需求配置子字段且该参数必填。 - 当匹配逻辑为num_greater、num_less、num_equal、num_not_equal时，子字段必须为空 - 当子字段不为空时，最大长度不超过2048
	Index *string `json:"index,omitempty"`

	// 条件列表匹配逻辑。 如果字段类型category是url、user-agent或者referer， 匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 len_greater、 len_less、len_equal或者len_not_equal 如果字段类型category是ip、ipv6或method，匹配逻辑可以为： equal、not_equal 如果字段类型category是request_line或者request, 匹配逻辑可以为： len_greater、len_less、len_equal或者len_not_equal 如果字段类型category是params、cookie或者header, 匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 len_greater、 len_less、len_equal、len_not_equal、num_greater、num_less、num_equal、num_not_equal、exist或者not_exist
	LogicOperation string `json:"logic_operation"`

	// 条件列表逻辑匹配内容。 - 当匹配逻辑为exist或not_exist时，contents必须为空，其他情况下contents必填且长度不超过2048 - 当匹配逻辑包含\"len\"时，contents必须为0~65535的整数；当匹配逻辑包含\"num\"时，contents必须为0~512的整数 - 当category为method时, contents必须是1-64位大写字母
	Contents []string `json:"contents"`
}

func (o WafCustomCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafCustomCondition struct{}"
	}

	return strings.Join([]string{"WafCustomCondition", string(data)}, " ")
}

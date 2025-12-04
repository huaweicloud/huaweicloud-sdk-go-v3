package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CcCondition struct {

	// **参数解释：** 字段类型 **约束限制：** 不涉及 **取值范围：**  - url  - ip  - ipv6  - asn  - params  - cookie  - referer  - user-agent  - header  - response_code  - response_header  - response_body  - request_body  - method  - tls_fingerprint  - tls_ja3_fingerprint  **默认取值：** 不涉及
	Category CcConditionCategory `json:"category"`

	// 条件列表匹配逻辑。   -  如果字段类型category是url， 匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 contain_any、 not_contain_all、 equal_any、 not_equal_all、 equal_any、 not_equal_all、 prefix_any、 not_prefix_all、 suffix_any、 not_suffix_all、 len_greater、 len_less、len_equal或者len_not_equal   - 如果字段类型category是ip或者ipv6，匹配逻辑可以为： equal、not_equal、equal_any或者not_equal_all   - 如果字段类型category是params、cookie或者header, 匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 contain_any、 not_contain_all、 equal_any、not_equal_all、 equal_any、 not_equal_all、 prefix_any、 not_prefix_all、 suffix_any、 not_suffix_all、 len_greater、 len_less、len_equal、len_not_equal、、num_greater、num_less、num_equal、num_not_equal、exist或者not_exist
	LogicOperation string `json:"logic_operation"`

	// 条件列表逻辑匹配内容。当logic_operation参数不以any或者all结尾时，需要传该参数。
	Contents *[]string `json:"contents,omitempty"`

	// 引用表id。当logic_operation参数以any或者all结尾时，需要传该参数。此外，引用表类型要与category类型保持一致。
	ValueListId *string `json:"value_list_id,omitempty"`

	// 子字段，当字段类型（category）选择“params”、“cookie”、“header”时，请根据实际需求配置子字段且该参数必填。
	Index *string `json:"index,omitempty"`
}

func (o CcCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CcCondition struct{}"
	}

	return strings.Join([]string{"CcCondition", string(data)}, " ")
}

type CcConditionCategory struct {
	value string
}

type CcConditionCategoryEnum struct {
	URL                 CcConditionCategory
	IP                  CcConditionCategory
	IPV6                CcConditionCategory
	ASN                 CcConditionCategory
	PARAMS              CcConditionCategory
	COOKIE              CcConditionCategory
	REFERER             CcConditionCategory
	USER_AGENT          CcConditionCategory
	HEADER              CcConditionCategory
	RESPONSE_CODE       CcConditionCategory
	RESPONSE_HEADER     CcConditionCategory
	RESPONSE_BODY       CcConditionCategory
	REQUEST_BODY        CcConditionCategory
	METHOD              CcConditionCategory
	TLS_FINGERPRINT     CcConditionCategory
	TLS_JA3_FINGERPRINT CcConditionCategory
}

func GetCcConditionCategoryEnum() CcConditionCategoryEnum {
	return CcConditionCategoryEnum{
		URL: CcConditionCategory{
			value: "url",
		},
		IP: CcConditionCategory{
			value: "ip",
		},
		IPV6: CcConditionCategory{
			value: "ipv6",
		},
		ASN: CcConditionCategory{
			value: "asn",
		},
		PARAMS: CcConditionCategory{
			value: "params",
		},
		COOKIE: CcConditionCategory{
			value: "cookie",
		},
		REFERER: CcConditionCategory{
			value: "referer",
		},
		USER_AGENT: CcConditionCategory{
			value: "user-agent",
		},
		HEADER: CcConditionCategory{
			value: "header",
		},
		RESPONSE_CODE: CcConditionCategory{
			value: "response_code",
		},
		RESPONSE_HEADER: CcConditionCategory{
			value: "response_header",
		},
		RESPONSE_BODY: CcConditionCategory{
			value: "response_body",
		},
		REQUEST_BODY: CcConditionCategory{
			value: "request_body",
		},
		METHOD: CcConditionCategory{
			value: "method",
		},
		TLS_FINGERPRINT: CcConditionCategory{
			value: "tls_fingerprint",
		},
		TLS_JA3_FINGERPRINT: CcConditionCategory{
			value: "tls_ja3_fingerprint",
		},
	}
}

func (c CcConditionCategory) Value() string {
	return c.value
}

func (c CcConditionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CcConditionCategory) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

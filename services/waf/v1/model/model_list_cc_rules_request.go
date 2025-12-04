package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCcRulesRequest Request Object
type ListCcRulesRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 策略id（策略id从查询防护策略列表接口获取）
	PolicyId string `json:"policy_id"`

	// 偏移量，表示查询该偏移量之后的记录。
	Offset int32 `json:"offset"`

	// 查询返回记录的数量限制。
	Limit int32 `json:"limit"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则的开启状态，1表示开启，0表示关闭
	Status *int32 `json:"status,omitempty"`

	// **参数解释：** 防护动作 **取值范围：** - block: 拦截，表示超过“限速频率”将直接拦截。 - log：仅记录，表示超过“限速频率”将只记录不拦截。 - captcha：表示超过“限速频率”后弹出验证码，进行人机验证，完成验证后，请求将不受访问限制。人机验证目前支持英文。 - dynamic_block：上一个限速周期内，请求频率超过“限速频率”将被拦截，那么在下一个限速周期内，请求频率超过“放行频率”将被拦截。 - advanced_captcha：高阶人机验证，表示超过“限速频率”后弹出验证码，进行人机验证。 - js_challenge：要求客户端完成一段脚本的执行或验证，从而验证请求来源的合法性。
	Category *string `json:"category,omitempty"`

	// **参数解释：** 限速模式标识，用于指定区分单个Web访问者的判断依据 **约束限制：** 不涉及 **取值范围：**  - ip：IP限速，根据IP区分单个Web访问者  - cookie：用户限速，根据Cookie键值区分单个Web访问者  - header：用户限速，根据Header区分单个Web访问者  - other：根据Referer（自定义请求访问的来源）字段区分单个Web访问者  - policy：策略限速  - domain：域名限速  - url：url限速 **默认取值：** 不涉及
	TagType *ListCcRulesRequestTagType `json:"tag_type,omitempty"`
}

func (o ListCcRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCcRulesRequest struct{}"
	}

	return strings.Join([]string{"ListCcRulesRequest", string(data)}, " ")
}

type ListCcRulesRequestTagType struct {
	value string
}

type ListCcRulesRequestTagTypeEnum struct {
	IP     ListCcRulesRequestTagType
	COOKIE ListCcRulesRequestTagType
	HEADER ListCcRulesRequestTagType
	OTHER  ListCcRulesRequestTagType
	POLICY ListCcRulesRequestTagType
	DOMAIN ListCcRulesRequestTagType
	URL    ListCcRulesRequestTagType
}

func GetListCcRulesRequestTagTypeEnum() ListCcRulesRequestTagTypeEnum {
	return ListCcRulesRequestTagTypeEnum{
		IP: ListCcRulesRequestTagType{
			value: "ip",
		},
		COOKIE: ListCcRulesRequestTagType{
			value: "cookie",
		},
		HEADER: ListCcRulesRequestTagType{
			value: "header",
		},
		OTHER: ListCcRulesRequestTagType{
			value: "other",
		},
		POLICY: ListCcRulesRequestTagType{
			value: "policy",
		},
		DOMAIN: ListCcRulesRequestTagType{
			value: "domain",
		},
		URL: ListCcRulesRequestTagType{
			value: "url",
		},
	}
}

func (c ListCcRulesRequestTagType) Value() string {
	return c.value
}

func (c ListCcRulesRequestTagType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCcRulesRequestTagType) UnmarshalJSON(b []byte) error {
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

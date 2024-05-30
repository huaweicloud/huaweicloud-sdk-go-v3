package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCustomerV2Req struct {

	// 客户的华为云账号名。 如果为空，随机生成。 不能以“op_”或“shadow_”开头且不能全为数字。 校验长度（5到32位）和规则^([a-zA-Z_-]([a-zA-Z0-9_-])*)$。 此参数不携带或携带值为空串或携带值为null时，随机生成。
	DomainName *string `json:"domain_name,omitempty"`

	// 手机号。 目前系统只支持中国的手机号。 示例：13XXXXXXXXX 此参数不携带或携带值为null时，不被赋值；携带值为空串时，赋值为空串。
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// 验证码。 请调用“发送验证码”接口获取。 如果手机号不存在，则不需要输入验证码。 此参数不携带或携带值为null时，不做处理；不支持携带值为空串。
	VerificationCode *string `json:"verification_code,omitempty"`

	// 伙伴销售平台的用户唯一标识，该标识的具体值由伙伴分配。
	XaccountId string `json:"xaccount_id"`

	// 华为分给合作伙伴的平台标识。 该标识的具体值由华为分配。获取方法请参见如何获取xaccountType的取值。
	XaccountType string `json:"xaccount_type"`

	// 密码规则如下： 至少包含以下四种字符中的两种： 大写字母、小写字母、数字、特殊字符；不能和账号名或倒序的账号名相同；不能包含手机号。 如果为空，用户没有密码，则不能直接在华为云登录，只能通过伙伴系统SSO方式跳转到华为云。 此参数不携带或携带值为null时，密码随机生成；不支持携带值为空串。
	Password *string `json:"password,omitempty"`

	// 是否关闭营销消息的发送。 true：关闭false：不关闭（默认） 此参数不携带或携带值为空串或携带值为null时，赋值为false。
	IsCloseMarketMs *string `json:"is_close_market_ms,omitempty"`

	// 合作类型。 1：顾问销售。 不传递或传递非1的值，默认会创建成代售模式的客户。
	CooperationType *string `json:"cooperation_type,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。 如果需要创建云经销商的子客户，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`

	// 是否返回子客户的关联结果。 true：返回子客户和伙伴的关联结果false：不返回子客户和伙伴的关联结果 默认值为false。 此参数不携带或携带值为空串或携带值为null时，赋值为false。
	IncludeAssociationResult *bool `json:"include_association_result,omitempty"`
}

func (o CreateCustomerV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomerV2Req struct{}"
	}

	return strings.Join([]string{"CreateCustomerV2Req", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerInformation struct {

	// 实名认证名称。
	Customer *string `json:"customer,omitempty"`

	// 客户登录名称（如果客户创建了子用户，此处返回主账号登录名称）。
	AccountName string `json:"account_name"`

	// 客户账号ID。
	CustomerId string `json:"customer_id"`

	// 客户和伙伴关联时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”，其中，HH范围是0～23，mm和ss范围是0～59。
	AssociatedOn *string `json:"associated_on,omitempty"`

	// 关联类型： 1：顾问销售2：代售
	AssociationType *string `json:"association_type,omitempty"`

	// 标签。
	Label *string `json:"label,omitempty"`

	// 客户电话号码。
	Telephone *string `json:"telephone,omitempty"`

	// 实名认证状态： -1：未实名认证0：实名认证审核中1：实名认证不通过2：已实名认证
	VerifiedStatus *string `json:"verified_status,omitempty"`

	// 国家码，电话号码的国家码前缀。 例如：中国 0086。
	CountryCode *string `json:"country_code,omitempty"`

	// 客户类型： -1：无类型0：个人1：企业 客户刚注册的时候，没有具体的客户类型，为“-1：无类型”，客户可以在账号中心通过设置客户类型或者在实名认证的时候，选择对应的企业/个人实名认证来决定自己的类型。
	CustomerType *int32 `json:"customer_type,omitempty"`

	// 是否冻结： 0：否1：客户账号冻结2：客户账号和资源冻结 该字段预留。
	IsFrozen *int32 `json:"is_frozen,omitempty"`

	// 该客户对应的客户经理信息，目前只支持1个，具体参见表2。
	AccountManagers *[]AccountManager `json:"account_managers,omitempty"`

	// 伙伴销售平台的用户唯一标识，该标识的具体值由伙伴分配。
	XaccountId *string `json:"xaccount_id,omitempty"`

	// 华为分配给合作伙伴的平台标识。 该标识的具体值由华为分配。获取方法请参见如何获取xaccountType的取值。
	XaccountType *string `json:"xaccount_type,omitempty"`

	// 客户等级。具体等级体系和权益请参见客户等级体系。 V0V1V2V3V4V5
	CustomerLevel *string `json:"customer_level,omitempty"`

	// |参数名称：客户邮箱| |参数的约束及描述：该参数为字符串 范围限制:0-256。|
	Email *string `json:"email,omitempty"`
}

func (o CustomerInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerInformation struct{}"
	}

	return strings.Join([]string{"CustomerInformation", string(data)}, " ")
}

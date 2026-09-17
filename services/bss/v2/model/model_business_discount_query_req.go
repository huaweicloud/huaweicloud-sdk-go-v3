package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BusinessDiscountQueryReq 查询商务折扣信息请求体
type BusinessDiscountQueryReq struct {

	// 报价项类型，必填，PRODUCT_ITEM（产品报价项）/ CATEGORY_ITEM（分类报价项）
	QuotingItemType BusinessDiscountQueryReqQuotingItemType `json:"quoting_item_type"`

	// 云服务类型编码列表，非必填，大小写不敏感，数组范围限制:0-100，字符长度限制1-64。此参数不携带或携带值为空列表或携带值为null时，不作为筛选条件。
	CloudServiceTypes *[]string `json:"cloud_service_types,omitempty"`

	// 计费模式列表，非必填，大小写不敏感，数组范围限制:0-20，字符长度限制1-64。此参数不携带或携带值为空列表或携带值为null时，不作为筛选条件。
	ChargingModes *[]string `json:"charging_modes,omitempty"`

	// 运营站点编码，非必填，大小写不敏感，字符长度限制1-64。此参数不携带或携带值为null时，不作为筛选条件。
	SiteCode *string `json:"site_code,omitempty"`

	// 分页偏移量，非必填，取值范围0-2147483647，默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数，非必填，取值范围1-1000，默认值20
	Limit *int32 `json:"limit,omitempty"`
}

func (o BusinessDiscountQueryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessDiscountQueryReq struct{}"
	}

	return strings.Join([]string{"BusinessDiscountQueryReq", string(data)}, " ")
}

type BusinessDiscountQueryReqQuotingItemType struct {
	value string
}

type BusinessDiscountQueryReqQuotingItemTypeEnum struct {
	PRODUCT_ITEM  BusinessDiscountQueryReqQuotingItemType
	CATEGORY_ITEM BusinessDiscountQueryReqQuotingItemType
}

func GetBusinessDiscountQueryReqQuotingItemTypeEnum() BusinessDiscountQueryReqQuotingItemTypeEnum {
	return BusinessDiscountQueryReqQuotingItemTypeEnum{
		PRODUCT_ITEM: BusinessDiscountQueryReqQuotingItemType{
			value: "PRODUCT_ITEM",
		},
		CATEGORY_ITEM: BusinessDiscountQueryReqQuotingItemType{
			value: "CATEGORY_ITEM",
		},
	}
}

func (c BusinessDiscountQueryReqQuotingItemType) Value() string {
	return c.value
}

func (c BusinessDiscountQueryReqQuotingItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BusinessDiscountQueryReqQuotingItemType) UnmarshalJSON(b []byte) error {
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

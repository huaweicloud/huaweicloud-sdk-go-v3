package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ProductDetailInfo 商品详细信息
type ProductDetailInfo struct {

	// 商品ID
	ProductId *string `json:"product_id,omitempty"`

	// 商品名称
	Name *string `json:"name,omitempty"`

	// 商品描述
	Description *string `json:"description,omitempty"`

	// 标签。单个标签16字节，多个用逗号分隔，最多50个。
	Tags *[]string `json:"tags,omitempty"`

	Cover *ProductCoverDetailInfo `json:"cover,omitempty"`

	// 文本列表
	TextList *[]ProductTextInfo `json:"text_list,omitempty"`

	// 素材资产列表
	AssetList *[]ProductMediaDetailInfo `json:"asset_list,omitempty"`

	// 商品创建时间，格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 商品更新时间，格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`

	// 商品状态枚举 * ACTIVED：已激活 * UNACTIVED：未激活 * BLOCK: 被冻结，商品不可用 * DELETED：已删除
	State *ProductDetailInfoState `json:"state,omitempty"`
}

func (o ProductDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductDetailInfo struct{}"
	}

	return strings.Join([]string{"ProductDetailInfo", string(data)}, " ")
}

type ProductDetailInfoState struct {
	value string
}

type ProductDetailInfoStateEnum struct {
	ACTIVED   ProductDetailInfoState
	UNACTIVED ProductDetailInfoState
	BLOCK     ProductDetailInfoState
	DELETED   ProductDetailInfoState
}

func GetProductDetailInfoStateEnum() ProductDetailInfoStateEnum {
	return ProductDetailInfoStateEnum{
		ACTIVED: ProductDetailInfoState{
			value: "ACTIVED",
		},
		UNACTIVED: ProductDetailInfoState{
			value: "UNACTIVED",
		},
		BLOCK: ProductDetailInfoState{
			value: "BLOCK",
		},
		DELETED: ProductDetailInfoState{
			value: "DELETED",
		},
	}
}

func (c ProductDetailInfoState) Value() string {
	return c.value
}

func (c ProductDetailInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProductDetailInfoState) UnmarshalJSON(b []byte) error {
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

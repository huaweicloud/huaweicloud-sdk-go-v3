package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ProductMediaInfo 商品素材信息
type ProductMediaInfo struct {

	// 资产ID
	AssetId *string `json:"asset_id,omitempty"`

	// 资产类型 * IMAGE：图片 * VIDEO：视频 * AUDIO：音频
	AssetType *ProductMediaInfoAssetType `json:"asset_type,omitempty"`

	// **参数解释**： 资产次序。不设置或者0表示按照加入时间先后排序。业务上将次序最靠前的图片设置为商品封面。
	Order *int32 `json:"order,omitempty"`
}

func (o ProductMediaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductMediaInfo struct{}"
	}

	return strings.Join([]string{"ProductMediaInfo", string(data)}, " ")
}

type ProductMediaInfoAssetType struct {
	value string
}

type ProductMediaInfoAssetTypeEnum struct {
	IMAGE ProductMediaInfoAssetType
	VIDEO ProductMediaInfoAssetType
	AUDIO ProductMediaInfoAssetType
}

func GetProductMediaInfoAssetTypeEnum() ProductMediaInfoAssetTypeEnum {
	return ProductMediaInfoAssetTypeEnum{
		IMAGE: ProductMediaInfoAssetType{
			value: "IMAGE",
		},
		VIDEO: ProductMediaInfoAssetType{
			value: "VIDEO",
		},
		AUDIO: ProductMediaInfoAssetType{
			value: "AUDIO",
		},
	}
}

func (c ProductMediaInfoAssetType) Value() string {
	return c.value
}

func (c ProductMediaInfoAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProductMediaInfoAssetType) UnmarshalJSON(b []byte) error {
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

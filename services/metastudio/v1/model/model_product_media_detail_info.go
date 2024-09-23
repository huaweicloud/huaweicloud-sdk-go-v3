package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProductMediaDetailInfo struct {

	// 资产ID
	AssetId *string `json:"asset_id,omitempty"`

	// 资产类型 * IMAGE：图片 * VIDEO：视频 * AUDIO：音频
	AssetType *ProductMediaDetailInfoAssetType `json:"asset_type,omitempty"`

	// **参数解释**： 资产次序。不设置或者0表示按照加入时间先后排序。业务上将次序最靠前的图片设置为商品封面。
	Order *int32 `json:"order,omitempty"`

	// 资产名称。
	AssetName *string `json:"asset_name,omitempty"`

	// 资产状态。
	AssetState *string `json:"asset_state,omitempty"`

	// 封面图片路径。
	CoverUrl *string `json:"cover_url,omitempty"`

	// 缩略图路径。
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`

	// 缩略图路径。
	MainUrl *string `json:"main_url,omitempty"`
}

func (o ProductMediaDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductMediaDetailInfo struct{}"
	}

	return strings.Join([]string{"ProductMediaDetailInfo", string(data)}, " ")
}

type ProductMediaDetailInfoAssetType struct {
	value string
}

type ProductMediaDetailInfoAssetTypeEnum struct {
	IMAGE ProductMediaDetailInfoAssetType
	VIDEO ProductMediaDetailInfoAssetType
	AUDIO ProductMediaDetailInfoAssetType
}

func GetProductMediaDetailInfoAssetTypeEnum() ProductMediaDetailInfoAssetTypeEnum {
	return ProductMediaDetailInfoAssetTypeEnum{
		IMAGE: ProductMediaDetailInfoAssetType{
			value: "IMAGE",
		},
		VIDEO: ProductMediaDetailInfoAssetType{
			value: "VIDEO",
		},
		AUDIO: ProductMediaDetailInfoAssetType{
			value: "AUDIO",
		},
	}
}

func (c ProductMediaDetailInfoAssetType) Value() string {
	return c.value
}

func (c ProductMediaDetailInfoAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProductMediaDetailInfoAssetType) UnmarshalJSON(b []byte) error {
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

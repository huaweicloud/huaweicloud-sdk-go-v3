package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StyleAssetItem 风格化素材数字资产信息
type StyleAssetItem struct {

	// 资产ID
	AssetId *string `json:"asset_id,omitempty"`

	// 资产类型 * ANIMATION：动作 * MATERIAL：素材
	AssetType *StyleAssetItemAssetType `json:"asset_type,omitempty"`

	// 封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`
}

func (o StyleAssetItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleAssetItem struct{}"
	}

	return strings.Join([]string{"StyleAssetItem", string(data)}, " ")
}

type StyleAssetItemAssetType struct {
	value string
}

type StyleAssetItemAssetTypeEnum struct {
	ANIMATION StyleAssetItemAssetType
	MATERIAL  StyleAssetItemAssetType
}

func GetStyleAssetItemAssetTypeEnum() StyleAssetItemAssetTypeEnum {
	return StyleAssetItemAssetTypeEnum{
		ANIMATION: StyleAssetItemAssetType{
			value: "ANIMATION",
		},
		MATERIAL: StyleAssetItemAssetType{
			value: "MATERIAL",
		},
	}
}

func (c StyleAssetItemAssetType) Value() string {
	return c.value
}

func (c StyleAssetItemAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StyleAssetItemAssetType) UnmarshalJSON(b []byte) error {
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

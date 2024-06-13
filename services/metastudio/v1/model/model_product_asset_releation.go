package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ProductAssetReleation 更新资产组合
type ProductAssetReleation struct {
	AssetItem *ProductMediaInfo `json:"asset_item,omitempty"`

	TextItem *ProductTextInfo `json:"text_item,omitempty"`

	// 添加或删除资产关联 - LINK ：将资产纳入管理 - UNLINK ：将资产移除管理
	Action *ProductAssetReleationAction `json:"action,omitempty"`
}

func (o ProductAssetReleation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductAssetReleation struct{}"
	}

	return strings.Join([]string{"ProductAssetReleation", string(data)}, " ")
}

type ProductAssetReleationAction struct {
	value string
}

type ProductAssetReleationActionEnum struct {
	LINK   ProductAssetReleationAction
	UNLINK ProductAssetReleationAction
}

func GetProductAssetReleationActionEnum() ProductAssetReleationActionEnum {
	return ProductAssetReleationActionEnum{
		LINK: ProductAssetReleationAction{
			value: "LINK",
		},
		UNLINK: ProductAssetReleationAction{
			value: "UNLINK",
		},
	}
}

func (c ProductAssetReleationAction) Value() string {
	return c.value
}

func (c ProductAssetReleationAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProductAssetReleationAction) UnmarshalJSON(b []byte) error {
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

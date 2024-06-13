package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateProductRequestBody 服务开通请求
type UpdateProductRequestBody struct {

	// 商品名称
	Name string `json:"name"`

	// 商品描述
	Description *string `json:"description,omitempty"`

	// 标签。单个标签16字节，多个用逗号分隔，最多50个。
	Tags *[]string `json:"tags,omitempty"`

	Cover *ProductCoverInfo `json:"cover,omitempty"`

	// 文本列表
	TextList *[]ProductTextInfo `json:"text_list,omitempty"`

	// 素材资产列表
	AssetList *[]ProductMediaInfo `json:"asset_list,omitempty"`

	// 商品状态枚举 * ACTIVED：已激活 * UNACTIVED：未激活
	State *UpdateProductRequestBodyState `json:"state,omitempty"`
}

func (o UpdateProductRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProductRequestBody", string(data)}, " ")
}

type UpdateProductRequestBodyState struct {
	value string
}

type UpdateProductRequestBodyStateEnum struct {
	ACTIVED   UpdateProductRequestBodyState
	UNACTIVED UpdateProductRequestBodyState
}

func GetUpdateProductRequestBodyStateEnum() UpdateProductRequestBodyStateEnum {
	return UpdateProductRequestBodyStateEnum{
		ACTIVED: UpdateProductRequestBodyState{
			value: "ACTIVED",
		},
		UNACTIVED: UpdateProductRequestBodyState{
			value: "UNACTIVED",
		},
	}
}

func (c UpdateProductRequestBodyState) Value() string {
	return c.value
}

func (c UpdateProductRequestBodyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateProductRequestBodyState) UnmarshalJSON(b []byte) error {
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

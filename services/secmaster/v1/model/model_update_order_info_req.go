package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateOrderInfoReq 云脑订单更新接口，包括包周期、按需场景，关联的变更场景
type UpdateOrderInfoReq struct {

	// 操作场景，包周期场景：PREPAID 或者 按需场景：POSTPAID
	Scene UpdateOrderInfoReqScene `json:"scene"`

	// 按需或者包周期场景下，要进行的操作类型 比如增减配额，规格升降配，按需转包周期 规格升级：UPGRADE，配额增加：ADDITION，配额减少：DECREASE，按需转包周期：POSTPAID_2_PREPAID 注：目前不支持规格降级，比如不支持从专业版降级为标准版或基础版
	OperateType UpdateOrderInfoReqOperateType `json:"operate_type"`

	// 促销折扣信息
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 计费标签
	TagList *[]SubscriptionTag `json:"tag_list,omitempty"`

	// 要进行变更的商品列表
	ProductList *[]UpdateProduct `json:"product_list,omitempty"`
}

func (o UpdateOrderInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrderInfoReq struct{}"
	}

	return strings.Join([]string{"UpdateOrderInfoReq", string(data)}, " ")
}

type UpdateOrderInfoReqScene struct {
	value string
}

type UpdateOrderInfoReqSceneEnum struct {
	PREPAID  UpdateOrderInfoReqScene
	POSTPAID UpdateOrderInfoReqScene
}

func GetUpdateOrderInfoReqSceneEnum() UpdateOrderInfoReqSceneEnum {
	return UpdateOrderInfoReqSceneEnum{
		PREPAID: UpdateOrderInfoReqScene{
			value: "PREPAID",
		},
		POSTPAID: UpdateOrderInfoReqScene{
			value: "POSTPAID",
		},
	}
}

func (c UpdateOrderInfoReqScene) Value() string {
	return c.value
}

func (c UpdateOrderInfoReqScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateOrderInfoReqScene) UnmarshalJSON(b []byte) error {
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

type UpdateOrderInfoReqOperateType struct {
	value string
}

type UpdateOrderInfoReqOperateTypeEnum struct {
	UPGRADE            UpdateOrderInfoReqOperateType
	ADDITION           UpdateOrderInfoReqOperateType
	DECREASE           UpdateOrderInfoReqOperateType
	POSTPAID_2_PREPAID UpdateOrderInfoReqOperateType
}

func GetUpdateOrderInfoReqOperateTypeEnum() UpdateOrderInfoReqOperateTypeEnum {
	return UpdateOrderInfoReqOperateTypeEnum{
		UPGRADE: UpdateOrderInfoReqOperateType{
			value: "UPGRADE",
		},
		ADDITION: UpdateOrderInfoReqOperateType{
			value: "ADDITION",
		},
		DECREASE: UpdateOrderInfoReqOperateType{
			value: "DECREASE",
		},
		POSTPAID_2_PREPAID: UpdateOrderInfoReqOperateType{
			value: "POSTPAID_2_PREPAID",
		},
	}
}

func (c UpdateOrderInfoReqOperateType) Value() string {
	return c.value
}

func (c UpdateOrderInfoReqOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateOrderInfoReqOperateType) UnmarshalJSON(b []byte) error {
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

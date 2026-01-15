package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResizeDesktopPoolReq 变更桌面池请求。
type ResizeDesktopPoolReq struct {

	// 目标规格产品ID。
	ProductId string `json:"product_id"`

	// 产品规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 是否支持开机状态下执行变更规格操作。固定传值STOP_DESKTOP，如果桌面处于开机状态，会先关机再变更规格。
	Mode *string `json:"mode,omitempty"`

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *ResizeDesktopPoolReqHandleType `json:"handle_type,omitempty"`

	// 桌面id
	DesktopIds *[]string `json:"desktop_ids,omitempty"`
}

func (o ResizeDesktopPoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopPoolReq struct{}"
	}

	return strings.Join([]string{"ResizeDesktopPoolReq", string(data)}, " ")
}

type ResizeDesktopPoolReqHandleType struct {
	value string
}

type ResizeDesktopPoolReqHandleTypeEnum struct {
	ONLY_FOR_EXPAND     ResizeDesktopPoolReqHandleType
	FOR_EXPAND_AND_IDLE ResizeDesktopPoolReqHandleType
	FOR_EXPAND_AND_ALL  ResizeDesktopPoolReqHandleType
}

func GetResizeDesktopPoolReqHandleTypeEnum() ResizeDesktopPoolReqHandleTypeEnum {
	return ResizeDesktopPoolReqHandleTypeEnum{
		ONLY_FOR_EXPAND: ResizeDesktopPoolReqHandleType{
			value: "ONLY_FOR_EXPAND",
		},
		FOR_EXPAND_AND_IDLE: ResizeDesktopPoolReqHandleType{
			value: "FOR_EXPAND_AND_IDLE",
		},
		FOR_EXPAND_AND_ALL: ResizeDesktopPoolReqHandleType{
			value: "FOR_EXPAND_AND_ALL",
		},
	}
}

func (c ResizeDesktopPoolReqHandleType) Value() string {
	return c.value
}

func (c ResizeDesktopPoolReqHandleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeDesktopPoolReqHandleType) UnmarshalJSON(b []byte) error {
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

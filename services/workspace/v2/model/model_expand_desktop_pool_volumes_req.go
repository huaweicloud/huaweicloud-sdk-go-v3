package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExpandDesktopPoolVolumesReq 扩容桌面池磁盘请求。
type ExpandDesktopPoolVolumesReq struct {

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *ExpandDesktopPoolVolumesReqHandleType `json:"handle_type,omitempty"`

	// 桌面id
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 扩容的桌面池磁盘列表。
	Volumes *[]ExpandDesktopPoolVolumesReqVolumes `json:"volumes,omitempty"`
}

func (o ExpandDesktopPoolVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopPoolVolumesReq struct{}"
	}

	return strings.Join([]string{"ExpandDesktopPoolVolumesReq", string(data)}, " ")
}

type ExpandDesktopPoolVolumesReqHandleType struct {
	value string
}

type ExpandDesktopPoolVolumesReqHandleTypeEnum struct {
	ONLY_FOR_EXPAND     ExpandDesktopPoolVolumesReqHandleType
	FOR_EXPAND_AND_IDLE ExpandDesktopPoolVolumesReqHandleType
	FOR_EXPAND_AND_ALL  ExpandDesktopPoolVolumesReqHandleType
}

func GetExpandDesktopPoolVolumesReqHandleTypeEnum() ExpandDesktopPoolVolumesReqHandleTypeEnum {
	return ExpandDesktopPoolVolumesReqHandleTypeEnum{
		ONLY_FOR_EXPAND: ExpandDesktopPoolVolumesReqHandleType{
			value: "ONLY_FOR_EXPAND",
		},
		FOR_EXPAND_AND_IDLE: ExpandDesktopPoolVolumesReqHandleType{
			value: "FOR_EXPAND_AND_IDLE",
		},
		FOR_EXPAND_AND_ALL: ExpandDesktopPoolVolumesReqHandleType{
			value: "FOR_EXPAND_AND_ALL",
		},
	}
}

func (c ExpandDesktopPoolVolumesReqHandleType) Value() string {
	return c.value
}

func (c ExpandDesktopPoolVolumesReqHandleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExpandDesktopPoolVolumesReqHandleType) UnmarshalJSON(b []byte) error {
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

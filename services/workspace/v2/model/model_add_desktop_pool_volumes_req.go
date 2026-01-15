package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddDesktopPoolVolumesReq 扩容桌面池请求。
type AddDesktopPoolVolumesReq struct {

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *AddDesktopPoolVolumesReqHandleType `json:"handle_type,omitempty"`

	// 桌面id
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 增加的磁盘列表。
	Volumes *[]VolumeAddInfo `json:"volumes,omitempty"`
}

func (o AddDesktopPoolVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopPoolVolumesReq struct{}"
	}

	return strings.Join([]string{"AddDesktopPoolVolumesReq", string(data)}, " ")
}

type AddDesktopPoolVolumesReqHandleType struct {
	value string
}

type AddDesktopPoolVolumesReqHandleTypeEnum struct {
	ONLY_FOR_EXPAND     AddDesktopPoolVolumesReqHandleType
	FOR_EXPAND_AND_IDLE AddDesktopPoolVolumesReqHandleType
	FOR_EXPAND_AND_ALL  AddDesktopPoolVolumesReqHandleType
}

func GetAddDesktopPoolVolumesReqHandleTypeEnum() AddDesktopPoolVolumesReqHandleTypeEnum {
	return AddDesktopPoolVolumesReqHandleTypeEnum{
		ONLY_FOR_EXPAND: AddDesktopPoolVolumesReqHandleType{
			value: "ONLY_FOR_EXPAND",
		},
		FOR_EXPAND_AND_IDLE: AddDesktopPoolVolumesReqHandleType{
			value: "FOR_EXPAND_AND_IDLE",
		},
		FOR_EXPAND_AND_ALL: AddDesktopPoolVolumesReqHandleType{
			value: "FOR_EXPAND_AND_ALL",
		},
	}
}

func (c AddDesktopPoolVolumesReqHandleType) Value() string {
	return c.value
}

func (c AddDesktopPoolVolumesReqHandleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddDesktopPoolVolumesReqHandleType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteDesktopPoolVolumesReq 删除桌面池磁盘请求。
type DeleteDesktopPoolVolumesReq struct {

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *DeleteDesktopPoolVolumesReqHandleType `json:"handle_type,omitempty"`

	// 桌面id
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 删除的桌面池磁盘列表。
	Volumes *[]DeleteDesktopPoolVolumesReqVolumes `json:"volumes,omitempty"`
}

func (o DeleteDesktopPoolVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopPoolVolumesReq struct{}"
	}

	return strings.Join([]string{"DeleteDesktopPoolVolumesReq", string(data)}, " ")
}

type DeleteDesktopPoolVolumesReqHandleType struct {
	value string
}

type DeleteDesktopPoolVolumesReqHandleTypeEnum struct {
	ONLY_FOR_EXPAND     DeleteDesktopPoolVolumesReqHandleType
	FOR_EXPAND_AND_IDLE DeleteDesktopPoolVolumesReqHandleType
	FOR_EXPAND_AND_ALL  DeleteDesktopPoolVolumesReqHandleType
}

func GetDeleteDesktopPoolVolumesReqHandleTypeEnum() DeleteDesktopPoolVolumesReqHandleTypeEnum {
	return DeleteDesktopPoolVolumesReqHandleTypeEnum{
		ONLY_FOR_EXPAND: DeleteDesktopPoolVolumesReqHandleType{
			value: "ONLY_FOR_EXPAND",
		},
		FOR_EXPAND_AND_IDLE: DeleteDesktopPoolVolumesReqHandleType{
			value: "FOR_EXPAND_AND_IDLE",
		},
		FOR_EXPAND_AND_ALL: DeleteDesktopPoolVolumesReqHandleType{
			value: "FOR_EXPAND_AND_ALL",
		},
	}
}

func (c DeleteDesktopPoolVolumesReqHandleType) Value() string {
	return c.value
}

func (c DeleteDesktopPoolVolumesReqHandleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDesktopPoolVolumesReqHandleType) UnmarshalJSON(b []byte) error {
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

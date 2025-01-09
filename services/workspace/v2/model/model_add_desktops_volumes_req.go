package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddDesktopsVolumesReq 增加磁盘请求。
type AddDesktopsVolumesReq struct {

	// 桌面类别  - DESKTOP：普通桌面。 - RENDER_DESKTOP：渲染桌面。
	DesktopType *AddDesktopsVolumesReqDesktopType `json:"desktop_type,omitempty"`

	// 新增磁盘参数。
	DesktopVolumes *[]AddDesktopVolumesReq `json:"desktop_volumes,omitempty"`
}

func (o AddDesktopsVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopsVolumesReq struct{}"
	}

	return strings.Join([]string{"AddDesktopsVolumesReq", string(data)}, " ")
}

type AddDesktopsVolumesReqDesktopType struct {
	value string
}

type AddDesktopsVolumesReqDesktopTypeEnum struct {
	DESKTOP        AddDesktopsVolumesReqDesktopType
	RENDER_DESKTOP AddDesktopsVolumesReqDesktopType
}

func GetAddDesktopsVolumesReqDesktopTypeEnum() AddDesktopsVolumesReqDesktopTypeEnum {
	return AddDesktopsVolumesReqDesktopTypeEnum{
		DESKTOP: AddDesktopsVolumesReqDesktopType{
			value: "DESKTOP",
		},
		RENDER_DESKTOP: AddDesktopsVolumesReqDesktopType{
			value: "RENDER_DESKTOP",
		},
	}
}

func (c AddDesktopsVolumesReqDesktopType) Value() string {
	return c.value
}

func (c AddDesktopsVolumesReqDesktopType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddDesktopsVolumesReqDesktopType) UnmarshalJSON(b []byte) error {
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

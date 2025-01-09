package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExpandDesktopsVolumesReq 扩容磁盘请求。
type ExpandDesktopsVolumesReq struct {

	// 桌面类别  - DESKTOP：普通桌面。 - RENDER-DESKTOP：渲染桌面。
	DesktopType *ExpandDesktopsVolumesReqDesktopType `json:"desktop_type,omitempty"`

	// 扩容磁盘参数。
	DesktopVolumesExpansion *[]ExpandVolumesReq `json:"desktop_volumes_expansion,omitempty"`
}

func (o ExpandDesktopsVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopsVolumesReq struct{}"
	}

	return strings.Join([]string{"ExpandDesktopsVolumesReq", string(data)}, " ")
}

type ExpandDesktopsVolumesReqDesktopType struct {
	value string
}

type ExpandDesktopsVolumesReqDesktopTypeEnum struct {
	DESKTOP        ExpandDesktopsVolumesReqDesktopType
	RENDER_DESKTOP ExpandDesktopsVolumesReqDesktopType
}

func GetExpandDesktopsVolumesReqDesktopTypeEnum() ExpandDesktopsVolumesReqDesktopTypeEnum {
	return ExpandDesktopsVolumesReqDesktopTypeEnum{
		DESKTOP: ExpandDesktopsVolumesReqDesktopType{
			value: "DESKTOP",
		},
		RENDER_DESKTOP: ExpandDesktopsVolumesReqDesktopType{
			value: "RENDER-DESKTOP",
		},
	}
}

func (c ExpandDesktopsVolumesReqDesktopType) Value() string {
	return c.value
}

func (c ExpandDesktopsVolumesReqDesktopType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExpandDesktopsVolumesReqDesktopType) UnmarshalJSON(b []byte) error {
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

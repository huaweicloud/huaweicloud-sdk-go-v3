package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateStorageModeTypeReq struct {

	// 降冷模式。 取值如下： - WHOLE：整个媒资粒度。 - ORIGIN：原文件粒度。
	StorageModeType *UpdateStorageModeTypeReqStorageModeType `json:"storage_mode_type,omitempty"`
}

func (o UpdateStorageModeTypeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStorageModeTypeReq struct{}"
	}

	return strings.Join([]string{"UpdateStorageModeTypeReq", string(data)}, " ")
}

type UpdateStorageModeTypeReqStorageModeType struct {
	value string
}

type UpdateStorageModeTypeReqStorageModeTypeEnum struct {
	WHOLE  UpdateStorageModeTypeReqStorageModeType
	ORIGIN UpdateStorageModeTypeReqStorageModeType
}

func GetUpdateStorageModeTypeReqStorageModeTypeEnum() UpdateStorageModeTypeReqStorageModeTypeEnum {
	return UpdateStorageModeTypeReqStorageModeTypeEnum{
		WHOLE: UpdateStorageModeTypeReqStorageModeType{
			value: "WHOLE",
		},
		ORIGIN: UpdateStorageModeTypeReqStorageModeType{
			value: "ORIGIN",
		},
	}
}

func (c UpdateStorageModeTypeReqStorageModeType) Value() string {
	return c.value
}

func (c UpdateStorageModeTypeReqStorageModeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateStorageModeTypeReqStorageModeType) UnmarshalJSON(b []byte) error {
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

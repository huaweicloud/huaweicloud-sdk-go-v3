package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RestoreDisasterReq 容灾回切请求参数。
type RestoreDisasterReq struct {

	// 容灾类型。
	DisasterType RestoreDisasterReqDisasterType `json:"disaster_type"`
}

func (o RestoreDisasterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDisasterReq struct{}"
	}

	return strings.Join([]string{"RestoreDisasterReq", string(data)}, " ")
}

type RestoreDisasterReqDisasterType struct {
	value string
}

type RestoreDisasterReqDisasterTypeEnum struct {
	STREAM RestoreDisasterReqDisasterType
	DORADO RestoreDisasterReqDisasterType
}

func GetRestoreDisasterReqDisasterTypeEnum() RestoreDisasterReqDisasterTypeEnum {
	return RestoreDisasterReqDisasterTypeEnum{
		STREAM: RestoreDisasterReqDisasterType{
			value: "stream",
		},
		DORADO: RestoreDisasterReqDisasterType{
			value: "dorado",
		},
	}
}

func (c RestoreDisasterReqDisasterType) Value() string {
	return c.value
}

func (c RestoreDisasterReqDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestoreDisasterReqDisasterType) UnmarshalJSON(b []byte) error {
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

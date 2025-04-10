package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReleaseDisasterReq 解除容灾请求参数。
type ReleaseDisasterReq struct {

	// 容灾类型。
	DisasterType ReleaseDisasterReqDisasterType `json:"disaster_type"`
}

func (o ReleaseDisasterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReleaseDisasterReq struct{}"
	}

	return strings.Join([]string{"ReleaseDisasterReq", string(data)}, " ")
}

type ReleaseDisasterReqDisasterType struct {
	value string
}

type ReleaseDisasterReqDisasterTypeEnum struct {
	STREAM ReleaseDisasterReqDisasterType
	DORADO ReleaseDisasterReqDisasterType
}

func GetReleaseDisasterReqDisasterTypeEnum() ReleaseDisasterReqDisasterTypeEnum {
	return ReleaseDisasterReqDisasterTypeEnum{
		STREAM: ReleaseDisasterReqDisasterType{
			value: "stream",
		},
		DORADO: ReleaseDisasterReqDisasterType{
			value: "dorado",
		},
	}
}

func (c ReleaseDisasterReqDisasterType) Value() string {
	return c.value
}

func (c ReleaseDisasterReqDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReleaseDisasterReqDisasterType) UnmarshalJSON(b []byte) error {
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

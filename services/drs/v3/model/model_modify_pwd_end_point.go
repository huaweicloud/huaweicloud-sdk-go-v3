package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 批量修改数据库密码请求列表
type ModifyPwdEndPoint struct {
	// 数据库密码

	DbPassword string `json:"db_password"`
	// 类型，so：源库；ta：目标库。

	EndPointType ModifyPwdEndPointEndPointType `json:"end_point_type"`
	// 任务id

	JobId string `json:"job_id"`

	Kerberos *KerberosVo `json:"kerberos,omitempty"`
}

func (o ModifyPwdEndPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPwdEndPoint struct{}"
	}

	return strings.Join([]string{"ModifyPwdEndPoint", string(data)}, " ")
}

type ModifyPwdEndPointEndPointType struct {
	value string
}

type ModifyPwdEndPointEndPointTypeEnum struct {
	SO ModifyPwdEndPointEndPointType
	TA ModifyPwdEndPointEndPointType
}

func GetModifyPwdEndPointEndPointTypeEnum() ModifyPwdEndPointEndPointTypeEnum {
	return ModifyPwdEndPointEndPointTypeEnum{
		SO: ModifyPwdEndPointEndPointType{
			value: "so",
		},
		TA: ModifyPwdEndPointEndPointType{
			value: "ta",
		},
	}
}

func (c ModifyPwdEndPointEndPointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyPwdEndPointEndPointType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

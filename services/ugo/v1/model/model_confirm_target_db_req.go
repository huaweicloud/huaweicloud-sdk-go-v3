package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 确认目标数据库版本的请求体。
type ConfirmTargetDbReq struct {

	// 评估项目ID。
	EvaluationProjectId string `json:"evaluation_project_id"`

	// 目标数据库类型。
	TargetDbType ConfirmTargetDbReqTargetDbType `json:"target_db_type"`

	// 目标数据库版本。（注意：该字段需要与 target_db_type 字段组合成有效的目标数据库类型与版本，当前支持以下组合： RDS for PostgreSQL-11； RDS for PostgreSQL-Enhanced Edition； RDS for MySQL-5.7; GaussDB(for MySQL) 8.0。)
	TargetDbVersion ConfirmTargetDbReqTargetDbVersion `json:"target_db_version"`
}

func (o ConfirmTargetDbReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmTargetDbReq struct{}"
	}

	return strings.Join([]string{"ConfirmTargetDbReq", string(data)}, " ")
}

type ConfirmTargetDbReqTargetDbType struct {
	value string
}

type ConfirmTargetDbReqTargetDbTypeEnum struct {
	RDS_FOR_MY_SQL      ConfirmTargetDbReqTargetDbType
	GAUSS_DB_FOR_MY_SQL ConfirmTargetDbReqTargetDbType
	RDS_FOR_POSTGRE_SQL ConfirmTargetDbReqTargetDbType
}

func GetConfirmTargetDbReqTargetDbTypeEnum() ConfirmTargetDbReqTargetDbTypeEnum {
	return ConfirmTargetDbReqTargetDbTypeEnum{
		RDS_FOR_MY_SQL: ConfirmTargetDbReqTargetDbType{
			value: "RDS for MySQL",
		},
		GAUSS_DB_FOR_MY_SQL: ConfirmTargetDbReqTargetDbType{
			value: "GaussDB(for MySQL)",
		},
		RDS_FOR_POSTGRE_SQL: ConfirmTargetDbReqTargetDbType{
			value: "RDS for PostgreSQL",
		},
	}
}

func (c ConfirmTargetDbReqTargetDbType) Value() string {
	return c.value
}

func (c ConfirmTargetDbReqTargetDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmTargetDbReqTargetDbType) UnmarshalJSON(b []byte) error {
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

type ConfirmTargetDbReqTargetDbVersion struct {
	value string
}

type ConfirmTargetDbReqTargetDbVersionEnum struct {
	E_5_7            ConfirmTargetDbReqTargetDbVersion
	E_8_0            ConfirmTargetDbReqTargetDbVersion
	E_11             ConfirmTargetDbReqTargetDbVersion
	ENHANCED_EDITION ConfirmTargetDbReqTargetDbVersion
}

func GetConfirmTargetDbReqTargetDbVersionEnum() ConfirmTargetDbReqTargetDbVersionEnum {
	return ConfirmTargetDbReqTargetDbVersionEnum{
		E_5_7: ConfirmTargetDbReqTargetDbVersion{
			value: "5.7",
		},
		E_8_0: ConfirmTargetDbReqTargetDbVersion{
			value: "8.0",
		},
		E_11: ConfirmTargetDbReqTargetDbVersion{
			value: "11",
		},
		ENHANCED_EDITION: ConfirmTargetDbReqTargetDbVersion{
			value: "Enhanced Edition",
		},
	}
}

func (c ConfirmTargetDbReqTargetDbVersion) Value() string {
	return c.value
}

func (c ConfirmTargetDbReqTargetDbVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmTargetDbReqTargetDbVersion) UnmarshalJSON(b []byte) error {
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

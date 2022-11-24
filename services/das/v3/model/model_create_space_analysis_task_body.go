package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSpaceAnalysisTaskBody struct {

	// 操作类型
	Operate CreateSpaceAnalysisTaskBodyOperate `json:"operate"`

	// 引擎类型
	DatastoreType CreateSpaceAnalysisTaskBodyDatastoreType `json:"datastore_type"`
}

func (o CreateSpaceAnalysisTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSpaceAnalysisTaskBody struct{}"
	}

	return strings.Join([]string{"CreateSpaceAnalysisTaskBody", string(data)}, " ")
}

type CreateSpaceAnalysisTaskBodyOperate struct {
	value string
}

type CreateSpaceAnalysisTaskBodyOperateEnum struct {
	REANALYSIS CreateSpaceAnalysisTaskBodyOperate
}

func GetCreateSpaceAnalysisTaskBodyOperateEnum() CreateSpaceAnalysisTaskBodyOperateEnum {
	return CreateSpaceAnalysisTaskBodyOperateEnum{
		REANALYSIS: CreateSpaceAnalysisTaskBodyOperate{
			value: "reanalysis",
		},
	}
}

func (c CreateSpaceAnalysisTaskBodyOperate) Value() string {
	return c.value
}

func (c CreateSpaceAnalysisTaskBodyOperate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSpaceAnalysisTaskBodyOperate) UnmarshalJSON(b []byte) error {
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

type CreateSpaceAnalysisTaskBodyDatastoreType struct {
	value string
}

type CreateSpaceAnalysisTaskBodyDatastoreTypeEnum struct {
	MY_SQL              CreateSpaceAnalysisTaskBodyDatastoreType
	GAUSS_DB_FOR_MY_SQL CreateSpaceAnalysisTaskBodyDatastoreType
}

func GetCreateSpaceAnalysisTaskBodyDatastoreTypeEnum() CreateSpaceAnalysisTaskBodyDatastoreTypeEnum {
	return CreateSpaceAnalysisTaskBodyDatastoreTypeEnum{
		MY_SQL: CreateSpaceAnalysisTaskBodyDatastoreType{
			value: "MySQL",
		},
		GAUSS_DB_FOR_MY_SQL: CreateSpaceAnalysisTaskBodyDatastoreType{
			value: "GaussDB(for MySQL)",
		},
	}
}

func (c CreateSpaceAnalysisTaskBodyDatastoreType) Value() string {
	return c.value
}

func (c CreateSpaceAnalysisTaskBodyDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSpaceAnalysisTaskBodyDatastoreType) UnmarshalJSON(b []byte) error {
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

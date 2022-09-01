package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询可用的AZ请求
type QueryAvailableNodeTypeReq struct {

	// 引擎类型
	EngineType string `json:"engine_type" xml:"engine_type"`

	// 迁移场景，migration-实时迁移,sync-实时同步,cloudDataGuard-实时灾备
	DbUseType QueryAvailableNodeTypeReqDbUseType `json:"db_use_type" xml:"db_use_type"`

	// 迁移方向，up ：入云 ，灾备场景时对应本云为备，down：出云，灾备场景时对应本云为主，non-dbs：自建
	JobDirection QueryAvailableNodeTypeReqJobDirection `json:"job_direction" xml:"job_direction"`

	// 规格类型。
	NodeType string `json:"node_type" xml:"node_type"`

	// 是否是双主灾备，不填默认为false
	MultiWrite *string `json:"multi_write,omitempty" xml:"multi_write"`
}

func (o QueryAvailableNodeTypeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAvailableNodeTypeReq struct{}"
	}

	return strings.Join([]string{"QueryAvailableNodeTypeReq", string(data)}, " ")
}

type QueryAvailableNodeTypeReqDbUseType struct {
	value string
}

type QueryAvailableNodeTypeReqDbUseTypeEnum struct {
	MIGRATION        QueryAvailableNodeTypeReqDbUseType
	SYNC             QueryAvailableNodeTypeReqDbUseType
	CLOUD_DATA_GUARD QueryAvailableNodeTypeReqDbUseType
}

func GetQueryAvailableNodeTypeReqDbUseTypeEnum() QueryAvailableNodeTypeReqDbUseTypeEnum {
	return QueryAvailableNodeTypeReqDbUseTypeEnum{
		MIGRATION: QueryAvailableNodeTypeReqDbUseType{
			value: "migration",
		},
		SYNC: QueryAvailableNodeTypeReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: QueryAvailableNodeTypeReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c QueryAvailableNodeTypeReqDbUseType) Value() string {
	return c.value
}

func (c QueryAvailableNodeTypeReqDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryAvailableNodeTypeReqDbUseType) UnmarshalJSON(b []byte) error {
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

type QueryAvailableNodeTypeReqJobDirection struct {
	value string
}

type QueryAvailableNodeTypeReqJobDirectionEnum struct {
	UP      QueryAvailableNodeTypeReqJobDirection
	DOWN    QueryAvailableNodeTypeReqJobDirection
	NON_DBS QueryAvailableNodeTypeReqJobDirection
}

func GetQueryAvailableNodeTypeReqJobDirectionEnum() QueryAvailableNodeTypeReqJobDirectionEnum {
	return QueryAvailableNodeTypeReqJobDirectionEnum{
		UP: QueryAvailableNodeTypeReqJobDirection{
			value: "up",
		},
		DOWN: QueryAvailableNodeTypeReqJobDirection{
			value: "down",
		},
		NON_DBS: QueryAvailableNodeTypeReqJobDirection{
			value: "non-dbs",
		},
	}
}

func (c QueryAvailableNodeTypeReqJobDirection) Value() string {
	return c.value
}

func (c QueryAvailableNodeTypeReqJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryAvailableNodeTypeReqJobDirection) UnmarshalJSON(b []byte) error {
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

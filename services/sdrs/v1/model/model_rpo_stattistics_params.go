package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// rpo超标记录
type RpoStattisticsParams struct {

	// 资源的RPO超标趋势记录id。
	Id string `json:"id"`

	// 资源的RPO超标趋势记录打点时间。默认格式为：\"yyyy-MM-dd HH:mm\"。
	PointTime string `json:"point_time"`

	// RPO超标的资源个数。
	ResourceNum int32 `json:"resource_num"`

	// RPO超标的资源类型。replication：表示查询复制对的RPO超标趋势记录。
	ResourceType RpoStattisticsParamsResourceType `json:"resource_type"`

	// 创建时间。默认格式为：\"yyyy-MM-dd HH:mm:ss.SSS\"，例：\"2019-04-01 12:00:00.000\"。
	CreatedAt string `json:"created_at"`

	// 更新时间。默认格式为：\"yyyy-MM-dd HH:mm:ss.SSS\"，例：\"2019-04-01 12:00:00.000\"。
	UpdatedAt string `json:"updated_at"`
}

func (o RpoStattisticsParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RpoStattisticsParams struct{}"
	}

	return strings.Join([]string{"RpoStattisticsParams", string(data)}, " ")
}

type RpoStattisticsParamsResourceType struct {
	value string
}

type RpoStattisticsParamsResourceTypeEnum struct {
	REPLICATION RpoStattisticsParamsResourceType
}

func GetRpoStattisticsParamsResourceTypeEnum() RpoStattisticsParamsResourceTypeEnum {
	return RpoStattisticsParamsResourceTypeEnum{
		REPLICATION: RpoStattisticsParamsResourceType{
			value: "replication",
		},
	}
}

func (c RpoStattisticsParamsResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RpoStattisticsParamsResourceType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListProtectionGroupsRequest struct {
	// 每次请求返回结果个数限制，取值范围为[0,1000]的正整数，默认值为1000。

	Limit *int32 `json:"limit,omitempty"`
	// 每次请求开始的下标，即偏移量，默认值为0。offset必须为数字，不能为负数。

	Offset *int32 `json:"offset,omitempty"`
	// 保护组状态。

	Status *string `json:"status,omitempty"`
	// 保护组的名称。支持模糊查询。

	Name *string `json:"name,omitempty"`
	// 查询场景类型。 status_abnormal：表示查询异常状态的保护组列表。 stop_protected：表示查询停止保护的保护组列表。 period_no_dr_drill：表示查询一段时间未做容灾演练的保护组，默认为三个月。 general或空时：该参数不生效。

	QueryType *ListProtectionGroupsRequestQueryType `json:"query_type,omitempty"`
	// 保护组的当前生产站点可用区。

	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListProtectionGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListProtectionGroupsRequest", string(data)}, " ")
}

type ListProtectionGroupsRequestQueryType struct {
	value string
}

type ListProtectionGroupsRequestQueryTypeEnum struct {
	STATUS_ABNORMAL    ListProtectionGroupsRequestQueryType
	STOP_PROTECTED     ListProtectionGroupsRequestQueryType
	PERIOD_NO_DR_DRILL ListProtectionGroupsRequestQueryType
	GENERAL            ListProtectionGroupsRequestQueryType
}

func GetListProtectionGroupsRequestQueryTypeEnum() ListProtectionGroupsRequestQueryTypeEnum {
	return ListProtectionGroupsRequestQueryTypeEnum{
		STATUS_ABNORMAL: ListProtectionGroupsRequestQueryType{
			value: "status_abnormal",
		},
		STOP_PROTECTED: ListProtectionGroupsRequestQueryType{
			value: " stop_protected",
		},
		PERIOD_NO_DR_DRILL: ListProtectionGroupsRequestQueryType{
			value: "period_no_dr_drill",
		},
		GENERAL: ListProtectionGroupsRequestQueryType{
			value: "general",
		},
	}
}

func (c ListProtectionGroupsRequestQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProtectionGroupsRequestQueryType) UnmarshalJSON(b []byte) error {
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

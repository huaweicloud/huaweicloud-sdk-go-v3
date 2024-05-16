package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgentSearchParam agent搜索参数。
type AgentSearchParam struct {

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// 环境id
	EnvId *int64 `json:"env_id,omitempty"`

	// 探针状态。
	Status *AgentSearchParamStatus `json:"status,omitempty"`

	// region英文名称。
	Region string `json:"region"`

	// 是否按照采集状态排序,默认不填则不按状态排序，填y则按照状态排序。
	OrderByStatus *string `json:"order_by_status,omitempty"`

	// 需要查询的页码，最小数为1。
	Page int32 `json:"page"`

	// 查询结果每页最多显示的条数。
	PageSize *int32 `json:"page_size,omitempty"`

	// 关键字。
	Keyword *string `json:"keyword,omitempty"`
}

func (o AgentSearchParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentSearchParam struct{}"
	}

	return strings.Join([]string{"AgentSearchParam", string(data)}, " ")
}

type AgentSearchParamStatus struct {
	value string
}

type AgentSearchParamStatusEnum struct {
	ONLINE  AgentSearchParamStatus
	DISABLE AgentSearchParamStatus
	OFFLINE AgentSearchParamStatus
}

func GetAgentSearchParamStatusEnum() AgentSearchParamStatusEnum {
	return AgentSearchParamStatusEnum{
		ONLINE: AgentSearchParamStatus{
			value: "online",
		},
		DISABLE: AgentSearchParamStatus{
			value: "disable",
		},
		OFFLINE: AgentSearchParamStatus{
			value: "offline",
		},
	}
}

func (c AgentSearchParamStatus) Value() string {
	return c.value
}

func (c AgentSearchParamStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgentSearchParamStatus) UnmarshalJSON(b []byte) error {
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

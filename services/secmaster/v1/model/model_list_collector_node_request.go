package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCollectorNodeRequest Request Object
type ListCollectorNodeRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 节点的健康状态 - NORMAL 正常 - ANOMALIES 异常 - FAULTS 故障 - LOST_CONTACT 失联  **约束限制** 不涉及 **取值范围**: - NORMAL - ANOMALIES - FAULTS - LOST_CONTACT  **默认值** 不涉及
	HealthStatus *ListCollectorNodeRequestHealthStatus `json:"health_status,omitempty"`

	// 地区
	Region *string `json:"region,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`

	// IP地址
	IpAddress *string `json:"ip_address,omitempty"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 目录排序 - asc 升序排列 - desc 降序排列  **约束限制** 不涉及 **取值范围**: - asc - desc  **默认值** 不涉及
	SortDir *ListCollectorNodeRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListCollectorNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorNodeRequest struct{}"
	}

	return strings.Join([]string{"ListCollectorNodeRequest", string(data)}, " ")
}

type ListCollectorNodeRequestHealthStatus struct {
	value string
}

type ListCollectorNodeRequestHealthStatusEnum struct {
	NORMAL       ListCollectorNodeRequestHealthStatus
	ANOMALIES    ListCollectorNodeRequestHealthStatus
	FAULTS       ListCollectorNodeRequestHealthStatus
	LOST_CONTACT ListCollectorNodeRequestHealthStatus
}

func GetListCollectorNodeRequestHealthStatusEnum() ListCollectorNodeRequestHealthStatusEnum {
	return ListCollectorNodeRequestHealthStatusEnum{
		NORMAL: ListCollectorNodeRequestHealthStatus{
			value: "NORMAL",
		},
		ANOMALIES: ListCollectorNodeRequestHealthStatus{
			value: "ANOMALIES",
		},
		FAULTS: ListCollectorNodeRequestHealthStatus{
			value: "FAULTS",
		},
		LOST_CONTACT: ListCollectorNodeRequestHealthStatus{
			value: "LOST_CONTACT",
		},
	}
}

func (c ListCollectorNodeRequestHealthStatus) Value() string {
	return c.value
}

func (c ListCollectorNodeRequestHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorNodeRequestHealthStatus) UnmarshalJSON(b []byte) error {
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

type ListCollectorNodeRequestSortDir struct {
	value string
}

type ListCollectorNodeRequestSortDirEnum struct {
	ASC  ListCollectorNodeRequestSortDir
	DESC ListCollectorNodeRequestSortDir
}

func GetListCollectorNodeRequestSortDirEnum() ListCollectorNodeRequestSortDirEnum {
	return ListCollectorNodeRequestSortDirEnum{
		ASC: ListCollectorNodeRequestSortDir{
			value: "asc",
		},
		DESC: ListCollectorNodeRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListCollectorNodeRequestSortDir) Value() string {
	return c.value
}

func (c ListCollectorNodeRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorNodeRequestSortDir) UnmarshalJSON(b []byte) error {
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

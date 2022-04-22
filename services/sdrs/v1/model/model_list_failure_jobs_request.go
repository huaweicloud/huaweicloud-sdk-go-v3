package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListFailureJobsRequest struct {

	// 失败任务状态。createFail：表示创建失败。deleteFail：表示删除失败。attachFail：表示挂载失败。detachFail：表示卸载失败。expandFail：表示扩容失败。resizeFail：表示变更规格失败。startFail：表示开启保护失败。stopFail：表示停止保护失败。reverseFail：表示切换失败。failoverFail：表示故障切换失败。reprotectFail : 表示重保护失败。
	FailureStatus *ListFailureJobsRequestFailureStatus `json:"failure_status,omitempty"`

	// 保护组资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 保护组ID。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 资源类型。server_groups：表示保护组。protected_instances：表示保护实例。replications：表示复制对。disaster_recovery_drills：表示容灾演练。
	ResourceType *ListFailureJobsRequestResourceType `json:"resource_type,omitempty"`

	// 每次请求返回结果个数限制。取值范围为[0,1000]的正整数，默认值为1000。
	Limit *int32 `json:"limit,omitempty"`

	// 每次请求开始的下标，即偏移量，默认值为0。offset必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListFailureJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFailureJobsRequest struct{}"
	}

	return strings.Join([]string{"ListFailureJobsRequest", string(data)}, " ")
}

type ListFailureJobsRequestFailureStatus struct {
	value string
}

type ListFailureJobsRequestFailureStatusEnum struct {
	CREATE_FAIL    ListFailureJobsRequestFailureStatus
	DELETE_FAIL    ListFailureJobsRequestFailureStatus
	ATTACH_FAIL    ListFailureJobsRequestFailureStatus
	DETACH_FAIL    ListFailureJobsRequestFailureStatus
	EXPAND_FAIL    ListFailureJobsRequestFailureStatus
	RESIZE_FAIL    ListFailureJobsRequestFailureStatus
	START_FAIL     ListFailureJobsRequestFailureStatus
	STOP_FAIL      ListFailureJobsRequestFailureStatus
	REVERSE_FAIL   ListFailureJobsRequestFailureStatus
	FAILOVER_FAIL  ListFailureJobsRequestFailureStatus
	REPROTECT_FAIL ListFailureJobsRequestFailureStatus
}

func GetListFailureJobsRequestFailureStatusEnum() ListFailureJobsRequestFailureStatusEnum {
	return ListFailureJobsRequestFailureStatusEnum{
		CREATE_FAIL: ListFailureJobsRequestFailureStatus{
			value: "createFail",
		},
		DELETE_FAIL: ListFailureJobsRequestFailureStatus{
			value: "deleteFail",
		},
		ATTACH_FAIL: ListFailureJobsRequestFailureStatus{
			value: "attachFail",
		},
		DETACH_FAIL: ListFailureJobsRequestFailureStatus{
			value: "detachFail",
		},
		EXPAND_FAIL: ListFailureJobsRequestFailureStatus{
			value: "expandFail",
		},
		RESIZE_FAIL: ListFailureJobsRequestFailureStatus{
			value: "resizeFail",
		},
		START_FAIL: ListFailureJobsRequestFailureStatus{
			value: "startFail",
		},
		STOP_FAIL: ListFailureJobsRequestFailureStatus{
			value: "stopFail",
		},
		REVERSE_FAIL: ListFailureJobsRequestFailureStatus{
			value: "reverseFail",
		},
		FAILOVER_FAIL: ListFailureJobsRequestFailureStatus{
			value: "failoverFail",
		},
		REPROTECT_FAIL: ListFailureJobsRequestFailureStatus{
			value: "reprotectFail",
		},
	}
}

func (c ListFailureJobsRequestFailureStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFailureJobsRequestFailureStatus) UnmarshalJSON(b []byte) error {
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

type ListFailureJobsRequestResourceType struct {
	value string
}

type ListFailureJobsRequestResourceTypeEnum struct {
	SERVER_GROUPS            ListFailureJobsRequestResourceType
	PROTECTED_INSTANCES      ListFailureJobsRequestResourceType
	REPLICATIONS             ListFailureJobsRequestResourceType
	DISASTER_RECOVERY_DRILLS ListFailureJobsRequestResourceType
}

func GetListFailureJobsRequestResourceTypeEnum() ListFailureJobsRequestResourceTypeEnum {
	return ListFailureJobsRequestResourceTypeEnum{
		SERVER_GROUPS: ListFailureJobsRequestResourceType{
			value: "server_groups",
		},
		PROTECTED_INSTANCES: ListFailureJobsRequestResourceType{
			value: "protected_instances",
		},
		REPLICATIONS: ListFailureJobsRequestResourceType{
			value: "replications",
		},
		DISASTER_RECOVERY_DRILLS: ListFailureJobsRequestResourceType{
			value: "disaster_recovery_drills",
		},
	}
}

func (c ListFailureJobsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFailureJobsRequestResourceType) UnmarshalJSON(b []byte) error {
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

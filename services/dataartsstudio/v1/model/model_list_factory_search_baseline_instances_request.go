package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFactorySearchBaselineInstancesRequest Request Object
type ListFactorySearchBaselineInstancesRequest struct {

	// DataArts Studio实例ID。
	InstanceId string `json:"instance_id"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	WorkspaceId string `json:"workspace_id"`

	// 基线任务名称。
	BaselineName *string `json:"baseline_name,omitempty"`

	// 责任人。
	OwnerName *string `json:"owner_name,omitempty"`

	// 基线类型，DAY天基线，HOUR小时基线，默认查询所有基线类型。
	Type *ListFactorySearchBaselineInstancesRequestType `json:"type,omitempty"`

	// 优先级，取值有1/2/3/4/5，默认查询所有优先级。当同时查询优先级为1/2/3时，样例如下：priority=1&priority=2&priority=3
	Priority *int32 `json:"priority,omitempty"`

	// 状态： - ERROR：异常 - SAFE：安全 - DANGEROUS：预警 - OVER：破线 默认查询所有状态。
	Status *ListFactorySearchBaselineInstancesRequestStatus `json:"status,omitempty"`

	// 完成状态：UNFINISH未完成，FINISH完成，默认查询所有状态。
	FinishStatus *ListFactorySearchBaselineInstancesRequestFinishStatus `json:"finish_status,omitempty"`

	// 开始时间戳，用于检索基线任务实例的承诺时间，单位毫秒。默认为当天0点0分0秒，当end_time有值时，该值不能为空。
	StartTime *int64 `json:"start_time,omitempty"`

	// 终止时间戳，用于检索基线任务实例的承诺时间，单位毫秒。默认为当天23点59分59秒，当start_time有值时，该值不能为空，最大查询范围为180天。
	EndTime *int64 `json:"end_time,omitempty"`

	// 排序规则，示例 start_time_ms asc，表示按照开始时间升序排序，有如下取值： - start_time_ms asc：按照开始时间升序。 - start_time_ms desc：按照开始时间降序。 - end_time_ms asc：按照结束时间升序。 - end_time_ms desc：按照结束时间降序。 默认不排序。
	OrderBy *string `json:"order_by,omitempty"`

	// 分页列表的页数，默认值为1。取值范围大于等于1。
	Offset *int32 `json:"offset,omitempty"`

	// 分页返回结果，指定每页最大记录数。范围[1,100] 默认值：10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFactorySearchBaselineInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactorySearchBaselineInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListFactorySearchBaselineInstancesRequest", string(data)}, " ")
}

type ListFactorySearchBaselineInstancesRequestType struct {
	value string
}

type ListFactorySearchBaselineInstancesRequestTypeEnum struct {
	DAY  ListFactorySearchBaselineInstancesRequestType
	HOUR ListFactorySearchBaselineInstancesRequestType
}

func GetListFactorySearchBaselineInstancesRequestTypeEnum() ListFactorySearchBaselineInstancesRequestTypeEnum {
	return ListFactorySearchBaselineInstancesRequestTypeEnum{
		DAY: ListFactorySearchBaselineInstancesRequestType{
			value: "DAY",
		},
		HOUR: ListFactorySearchBaselineInstancesRequestType{
			value: "HOUR",
		},
	}
}

func (c ListFactorySearchBaselineInstancesRequestType) Value() string {
	return c.value
}

func (c ListFactorySearchBaselineInstancesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFactorySearchBaselineInstancesRequestType) UnmarshalJSON(b []byte) error {
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

type ListFactorySearchBaselineInstancesRequestStatus struct {
	value string
}

type ListFactorySearchBaselineInstancesRequestStatusEnum struct {
	ERROR     ListFactorySearchBaselineInstancesRequestStatus
	SAFE      ListFactorySearchBaselineInstancesRequestStatus
	DANGEROUS ListFactorySearchBaselineInstancesRequestStatus
	OVER      ListFactorySearchBaselineInstancesRequestStatus
}

func GetListFactorySearchBaselineInstancesRequestStatusEnum() ListFactorySearchBaselineInstancesRequestStatusEnum {
	return ListFactorySearchBaselineInstancesRequestStatusEnum{
		ERROR: ListFactorySearchBaselineInstancesRequestStatus{
			value: "ERROR",
		},
		SAFE: ListFactorySearchBaselineInstancesRequestStatus{
			value: "SAFE",
		},
		DANGEROUS: ListFactorySearchBaselineInstancesRequestStatus{
			value: "DANGEROUS",
		},
		OVER: ListFactorySearchBaselineInstancesRequestStatus{
			value: "OVER",
		},
	}
}

func (c ListFactorySearchBaselineInstancesRequestStatus) Value() string {
	return c.value
}

func (c ListFactorySearchBaselineInstancesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFactorySearchBaselineInstancesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListFactorySearchBaselineInstancesRequestFinishStatus struct {
	value string
}

type ListFactorySearchBaselineInstancesRequestFinishStatusEnum struct {
	UNFINISH ListFactorySearchBaselineInstancesRequestFinishStatus
	FINISH   ListFactorySearchBaselineInstancesRequestFinishStatus
}

func GetListFactorySearchBaselineInstancesRequestFinishStatusEnum() ListFactorySearchBaselineInstancesRequestFinishStatusEnum {
	return ListFactorySearchBaselineInstancesRequestFinishStatusEnum{
		UNFINISH: ListFactorySearchBaselineInstancesRequestFinishStatus{
			value: "UNFINISH",
		},
		FINISH: ListFactorySearchBaselineInstancesRequestFinishStatus{
			value: "FINISH",
		},
	}
}

func (c ListFactorySearchBaselineInstancesRequestFinishStatus) Value() string {
	return c.value
}

func (c ListFactorySearchBaselineInstancesRequestFinishStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFactorySearchBaselineInstancesRequestFinishStatus) UnmarshalJSON(b []byte) error {
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

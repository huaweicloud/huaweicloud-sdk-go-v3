package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListConditionRequest Request Object
type ListConditionRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询
	Approver *string `json:"approver,omitempty"`

	// 业务状态
	Status *ListConditionRequestStatus `json:"status,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConditionRequest struct{}"
	}

	return strings.Join([]string{"ListConditionRequest", string(data)}, " ")
}

type ListConditionRequestStatus struct {
	value string
}

type ListConditionRequestStatusEnum struct {
	DRAFT              ListConditionRequestStatus
	PUBLISH_DEVELOPING ListConditionRequestStatus
	PUBLISHED          ListConditionRequestStatus
	OFFLINE_DEVELOPING ListConditionRequestStatus
	OFFLINE            ListConditionRequestStatus
	REJECT             ListConditionRequestStatus
}

func GetListConditionRequestStatusEnum() ListConditionRequestStatusEnum {
	return ListConditionRequestStatusEnum{
		DRAFT: ListConditionRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListConditionRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListConditionRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListConditionRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListConditionRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListConditionRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListConditionRequestStatus) Value() string {
	return c.value
}

func (c ListConditionRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConditionRequestStatus) UnmarshalJSON(b []byte) error {
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

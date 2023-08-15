package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchApprovalsRequest Request Object
type SearchApprovalsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 业务定义id
	BizId *int64 `json:"biz_id,omitempty"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询
	Approver *string `json:"approver,omitempty"`

	// 查询待审批，已审批
	ApprovalStatus *SearchApprovalsRequestApprovalStatus `json:"approval_status,omitempty"`

	// 查询待审批，已审批
	ApprovalStatusDetail *SearchApprovalsRequestApprovalStatusDetail `json:"approval_status_detail,omitempty"`

	// 业务审核类型
	ApprovalType *SearchApprovalsRequestApprovalType `json:"approval_type,omitempty"`

	// 按业务类型查询
	BizType *string `json:"biz_type,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchApprovalsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchApprovalsRequest struct{}"
	}

	return strings.Join([]string{"SearchApprovalsRequest", string(data)}, " ")
}

type SearchApprovalsRequestApprovalStatus struct {
	value string
}

type SearchApprovalsRequestApprovalStatusEnum struct {
	DEVELOPING SearchApprovalsRequestApprovalStatus
	FINISHED   SearchApprovalsRequestApprovalStatus
}

func GetSearchApprovalsRequestApprovalStatusEnum() SearchApprovalsRequestApprovalStatusEnum {
	return SearchApprovalsRequestApprovalStatusEnum{
		DEVELOPING: SearchApprovalsRequestApprovalStatus{
			value: "DEVELOPING",
		},
		FINISHED: SearchApprovalsRequestApprovalStatus{
			value: "FINISHED",
		},
	}
}

func (c SearchApprovalsRequestApprovalStatus) Value() string {
	return c.value
}

func (c SearchApprovalsRequestApprovalStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchApprovalsRequestApprovalStatus) UnmarshalJSON(b []byte) error {
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

type SearchApprovalsRequestApprovalStatusDetail struct {
	value string
}

type SearchApprovalsRequestApprovalStatusDetailEnum struct {
	DEVELOPING SearchApprovalsRequestApprovalStatusDetail
	APPROVED   SearchApprovalsRequestApprovalStatusDetail
	REJECT     SearchApprovalsRequestApprovalStatusDetail
}

func GetSearchApprovalsRequestApprovalStatusDetailEnum() SearchApprovalsRequestApprovalStatusDetailEnum {
	return SearchApprovalsRequestApprovalStatusDetailEnum{
		DEVELOPING: SearchApprovalsRequestApprovalStatusDetail{
			value: "DEVELOPING",
		},
		APPROVED: SearchApprovalsRequestApprovalStatusDetail{
			value: "APPROVED",
		},
		REJECT: SearchApprovalsRequestApprovalStatusDetail{
			value: "REJECT",
		},
	}
}

func (c SearchApprovalsRequestApprovalStatusDetail) Value() string {
	return c.value
}

func (c SearchApprovalsRequestApprovalStatusDetail) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchApprovalsRequestApprovalStatusDetail) UnmarshalJSON(b []byte) error {
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

type SearchApprovalsRequestApprovalType struct {
	value string
}

type SearchApprovalsRequestApprovalTypeEnum struct {
	PUBLISH SearchApprovalsRequestApprovalType
	OFFLINE SearchApprovalsRequestApprovalType
}

func GetSearchApprovalsRequestApprovalTypeEnum() SearchApprovalsRequestApprovalTypeEnum {
	return SearchApprovalsRequestApprovalTypeEnum{
		PUBLISH: SearchApprovalsRequestApprovalType{
			value: "PUBLISH",
		},
		OFFLINE: SearchApprovalsRequestApprovalType{
			value: "OFFLINE",
		},
	}
}

func (c SearchApprovalsRequestApprovalType) Value() string {
	return c.value
}

func (c SearchApprovalsRequestApprovalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchApprovalsRequestApprovalType) UnmarshalJSON(b []byte) error {
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

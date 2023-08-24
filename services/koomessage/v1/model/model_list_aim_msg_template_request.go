package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAimMsgTemplateRequest Request Object
type ListAimMsgTemplateRequest struct {

	// 偏移量。表示从偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 流程状态。
	FlowStatus *ListAimMsgTemplateRequestFlowStatus `json:"flow_status,omitempty"`

	// 模板ID。
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称。
	TemplateName *string `json:"template_name,omitempty"`

	// 模板类型。
	TemplateType *ListAimMsgTemplateRequestTemplateType `json:"template_type,omitempty"`

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`
}

func (o ListAimMsgTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimMsgTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListAimMsgTemplateRequest", string(data)}, " ")
}

type ListAimMsgTemplateRequestFlowStatus struct {
	value string
}

type ListAimMsgTemplateRequestFlowStatusEnum struct {
	ADOPTED   ListAimMsgTemplateRequestFlowStatus
	REVIEWING ListAimMsgTemplateRequestFlowStatus
	REJECT    ListAimMsgTemplateRequestFlowStatus
}

func GetListAimMsgTemplateRequestFlowStatusEnum() ListAimMsgTemplateRequestFlowStatusEnum {
	return ListAimMsgTemplateRequestFlowStatusEnum{
		ADOPTED: ListAimMsgTemplateRequestFlowStatus{
			value: "Adopted：通过",
		},
		REVIEWING: ListAimMsgTemplateRequestFlowStatus{
			value: "Reviewing：审核中",
		},
		REJECT: ListAimMsgTemplateRequestFlowStatus{
			value: "Reject：拒绝",
		},
	}
}

func (c ListAimMsgTemplateRequestFlowStatus) Value() string {
	return c.value
}

func (c ListAimMsgTemplateRequestFlowStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAimMsgTemplateRequestFlowStatus) UnmarshalJSON(b []byte) error {
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

type ListAimMsgTemplateRequestTemplateType struct {
	value string
}

type ListAimMsgTemplateRequestTemplateTypeEnum struct {
	PROMOTION_TYPE ListAimMsgTemplateRequestTemplateType
	NOTIFY_TYPE    ListAimMsgTemplateRequestTemplateType
}

func GetListAimMsgTemplateRequestTemplateTypeEnum() ListAimMsgTemplateRequestTemplateTypeEnum {
	return ListAimMsgTemplateRequestTemplateTypeEnum{
		PROMOTION_TYPE: ListAimMsgTemplateRequestTemplateType{
			value: "PROMOTION_TYPE：营销类",
		},
		NOTIFY_TYPE: ListAimMsgTemplateRequestTemplateType{
			value: "NOTIFY_TYPE：通知类",
		},
	}
}

func (c ListAimMsgTemplateRequestTemplateType) Value() string {
	return c.value
}

func (c ListAimMsgTemplateRequestTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAimMsgTemplateRequestTemplateType) UnmarshalJSON(b []byte) error {
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

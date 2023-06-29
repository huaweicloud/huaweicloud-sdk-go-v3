package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMyActionTemplateRequest Request Object
type ListMyActionTemplateRequest struct {

	// 模板前缀。
	Prefix *string `json:"prefix,omitempty"`

	// 第三方算子模板的注册状态。包括init_created, submit_approve, approved_nok, approved_ok, deprecate_approve, deprecated。init_created：已创建，submit_approve：提交审核，approved_nok：审核未通过，approved_ok：审核通过，deprecate_approve：提交禁用审核，deprecated：已禁用。
	Status *ListMyActionTemplateRequestStatus `json:"status,omitempty"`

	// 第三方算子模板的分类。默认分类为FileProcess,MediaProcess,ImageProcess,ContentReview,NotificationProcess,VoiceInteraction
	Category *ListMyActionTemplateRequestCategory `json:"category,omitempty"`

	// 查询的起始位置。start大于等于1，最大1000，不设置则取默认值1。
	Offset *string `json:"offset,omitempty"`

	// 分页查询，每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMyActionTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMyActionTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListMyActionTemplateRequest", string(data)}, " ")
}

type ListMyActionTemplateRequestStatus struct {
	value string
}

type ListMyActionTemplateRequestStatusEnum struct {
	INIT_CREATED      ListMyActionTemplateRequestStatus
	SUBMIT_APPROVE    ListMyActionTemplateRequestStatus
	APPROVED_NOK      ListMyActionTemplateRequestStatus
	APPROVED_OK       ListMyActionTemplateRequestStatus
	DEPRECATE_APPROVE ListMyActionTemplateRequestStatus
	DEPRECATED        ListMyActionTemplateRequestStatus
}

func GetListMyActionTemplateRequestStatusEnum() ListMyActionTemplateRequestStatusEnum {
	return ListMyActionTemplateRequestStatusEnum{
		INIT_CREATED: ListMyActionTemplateRequestStatus{
			value: "init_created",
		},
		SUBMIT_APPROVE: ListMyActionTemplateRequestStatus{
			value: " submit_approve",
		},
		APPROVED_NOK: ListMyActionTemplateRequestStatus{
			value: " approved_nok",
		},
		APPROVED_OK: ListMyActionTemplateRequestStatus{
			value: " approved_ok",
		},
		DEPRECATE_APPROVE: ListMyActionTemplateRequestStatus{
			value: " deprecate_approve",
		},
		DEPRECATED: ListMyActionTemplateRequestStatus{
			value: " deprecated",
		},
	}
}

func (c ListMyActionTemplateRequestStatus) Value() string {
	return c.value
}

func (c ListMyActionTemplateRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMyActionTemplateRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListMyActionTemplateRequestCategory struct {
	value string
}

type ListMyActionTemplateRequestCategoryEnum struct {
	FILE_PROCESS         ListMyActionTemplateRequestCategory
	MEDIA_PROCESS        ListMyActionTemplateRequestCategory
	IMAGE_PROCESS        ListMyActionTemplateRequestCategory
	CONTENT_REVIEW       ListMyActionTemplateRequestCategory
	NOTIFICATION_PROCESS ListMyActionTemplateRequestCategory
	VOICE_INTERACTION    ListMyActionTemplateRequestCategory
}

func GetListMyActionTemplateRequestCategoryEnum() ListMyActionTemplateRequestCategoryEnum {
	return ListMyActionTemplateRequestCategoryEnum{
		FILE_PROCESS: ListMyActionTemplateRequestCategory{
			value: "FileProcess",
		},
		MEDIA_PROCESS: ListMyActionTemplateRequestCategory{
			value: "MediaProcess",
		},
		IMAGE_PROCESS: ListMyActionTemplateRequestCategory{
			value: "ImageProcess",
		},
		CONTENT_REVIEW: ListMyActionTemplateRequestCategory{
			value: "ContentReview",
		},
		NOTIFICATION_PROCESS: ListMyActionTemplateRequestCategory{
			value: "NotificationProcess",
		},
		VOICE_INTERACTION: ListMyActionTemplateRequestCategory{
			value: "VoiceInteraction",
		},
	}
}

func (c ListMyActionTemplateRequestCategory) Value() string {
	return c.value
}

func (c ListMyActionTemplateRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMyActionTemplateRequestCategory) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfirmApprovalsRequest Request Object
type ConfirmApprovalsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 处理审批单结果类型。 枚举值：   - reject: 审批驳回   - resolve: 审批通过
	ActionId ConfirmApprovalsRequestActionId `json:"action-id"`

	Body *ApprovalInfoParam `json:"body,omitempty"`
}

func (o ConfirmApprovalsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmApprovalsRequest struct{}"
	}

	return strings.Join([]string{"ConfirmApprovalsRequest", string(data)}, " ")
}

type ConfirmApprovalsRequestActionId struct {
	value string
}

type ConfirmApprovalsRequestActionIdEnum struct {
	REJECT  ConfirmApprovalsRequestActionId
	RESOLVE ConfirmApprovalsRequestActionId
}

func GetConfirmApprovalsRequestActionIdEnum() ConfirmApprovalsRequestActionIdEnum {
	return ConfirmApprovalsRequestActionIdEnum{
		REJECT: ConfirmApprovalsRequestActionId{
			value: "reject",
		},
		RESOLVE: ConfirmApprovalsRequestActionId{
			value: "resolve",
		},
	}
}

func (c ConfirmApprovalsRequestActionId) Value() string {
	return c.value
}

func (c ConfirmApprovalsRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmApprovalsRequestActionId) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowOrganizationPolicyAssignmentDetailedStatusRequest Request Object
type ShowOrganizationPolicyAssignmentDetailedStatusRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 组织合规规则名称。
	OrganizationPolicyAssignmentName *string `json:"organization_policy_assignment_name,omitempty"`

	// 组织合规规则ID
	OrganizationPolicyAssignmentId *string `json:"organization_policy_assignment_id,omitempty"`

	// 成员帐号规则部署状态，区分大小写。
	Status *ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus `json:"status,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ShowOrganizationPolicyAssignmentDetailedStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationPolicyAssignmentDetailedStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationPolicyAssignmentDetailedStatusRequest", string(data)}, " ")
}

type ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus struct {
	value string
}

type ShowOrganizationPolicyAssignmentDetailedStatusRequestStatusEnum struct {
	CREATE_SUCCESSFUL  ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus
	CREATE_IN_PROGRESS ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus
	CREATE_FAILED      ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus
	DELETE_SUCCESSFUL  ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus
	DELETE_IN_PROGRESS ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus
	DELETE_FAILED      ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus
	UPDATE_SUCCESSFUL  ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus
	UPDATE_IN_PROGRESS ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus
	UPDATE_FAILED      ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus
}

func GetShowOrganizationPolicyAssignmentDetailedStatusRequestStatusEnum() ShowOrganizationPolicyAssignmentDetailedStatusRequestStatusEnum {
	return ShowOrganizationPolicyAssignmentDetailedStatusRequestStatusEnum{
		CREATE_SUCCESSFUL: ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus{
			value: "CREATE_SUCCESSFUL",
		},
		CREATE_IN_PROGRESS: ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus{
			value: "CREATE_IN_PROGRESS",
		},
		CREATE_FAILED: ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus{
			value: "CREATE_FAILED",
		},
		DELETE_SUCCESSFUL: ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus{
			value: "DELETE_SUCCESSFUL",
		},
		DELETE_IN_PROGRESS: ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus{
			value: "DELETE_IN_PROGRESS",
		},
		DELETE_FAILED: ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus{
			value: "DELETE_FAILED",
		},
		UPDATE_SUCCESSFUL: ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus{
			value: "UPDATE_SUCCESSFUL",
		},
		UPDATE_IN_PROGRESS: ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus{
			value: "UPDATE_IN_PROGRESS",
		},
		UPDATE_FAILED: ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus{
			value: "UPDATE_FAILED",
		},
	}
}

func (c ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus) Value() string {
	return c.value
}

func (c ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOrganizationPolicyAssignmentDetailedStatusRequestStatus) UnmarshalJSON(b []byte) error {
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

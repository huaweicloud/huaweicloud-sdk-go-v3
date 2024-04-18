package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowOrganizationConformancePackDetailedStatusesRequest Request Object
type ShowOrganizationConformancePackDetailedStatusesRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 合规规则包名称。
	ConformancePackName *string `json:"conformance_pack_name,omitempty"`

	// 组织合规规则包ID。
	OrganizationConformancePackId *string `json:"organization_conformance_pack_id,omitempty"`

	// 部署状态，区分大小写
	State *ShowOrganizationConformancePackDetailedStatusesRequestState `json:"state,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ShowOrganizationConformancePackDetailedStatusesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationConformancePackDetailedStatusesRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationConformancePackDetailedStatusesRequest", string(data)}, " ")
}

type ShowOrganizationConformancePackDetailedStatusesRequestState struct {
	value string
}

type ShowOrganizationConformancePackDetailedStatusesRequestStateEnum struct {
	CREATE_SUCCESSFUL  ShowOrganizationConformancePackDetailedStatusesRequestState
	CREATE_IN_PROGRESS ShowOrganizationConformancePackDetailedStatusesRequestState
	CREATE_FAILED      ShowOrganizationConformancePackDetailedStatusesRequestState
	DELETE_IN_PROGRESS ShowOrganizationConformancePackDetailedStatusesRequestState
	DELETE_FAILED      ShowOrganizationConformancePackDetailedStatusesRequestState
	UPDATE_SUCCESSFUL  ShowOrganizationConformancePackDetailedStatusesRequestState
	UPDATE_IN_PROGRESS ShowOrganizationConformancePackDetailedStatusesRequestState
	UPDATE_FAILED      ShowOrganizationConformancePackDetailedStatusesRequestState
}

func GetShowOrganizationConformancePackDetailedStatusesRequestStateEnum() ShowOrganizationConformancePackDetailedStatusesRequestStateEnum {
	return ShowOrganizationConformancePackDetailedStatusesRequestStateEnum{
		CREATE_SUCCESSFUL: ShowOrganizationConformancePackDetailedStatusesRequestState{
			value: "CREATE_SUCCESSFUL",
		},
		CREATE_IN_PROGRESS: ShowOrganizationConformancePackDetailedStatusesRequestState{
			value: "CREATE_IN_PROGRESS",
		},
		CREATE_FAILED: ShowOrganizationConformancePackDetailedStatusesRequestState{
			value: "CREATE_FAILED",
		},
		DELETE_IN_PROGRESS: ShowOrganizationConformancePackDetailedStatusesRequestState{
			value: "DELETE_IN_PROGRESS",
		},
		DELETE_FAILED: ShowOrganizationConformancePackDetailedStatusesRequestState{
			value: "DELETE_FAILED",
		},
		UPDATE_SUCCESSFUL: ShowOrganizationConformancePackDetailedStatusesRequestState{
			value: "UPDATE_SUCCESSFUL",
		},
		UPDATE_IN_PROGRESS: ShowOrganizationConformancePackDetailedStatusesRequestState{
			value: "UPDATE_IN_PROGRESS",
		},
		UPDATE_FAILED: ShowOrganizationConformancePackDetailedStatusesRequestState{
			value: "UPDATE_FAILED",
		},
	}
}

func (c ShowOrganizationConformancePackDetailedStatusesRequestState) Value() string {
	return c.value
}

func (c ShowOrganizationConformancePackDetailedStatusesRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOrganizationConformancePackDetailedStatusesRequestState) UnmarshalJSON(b []byte) error {
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

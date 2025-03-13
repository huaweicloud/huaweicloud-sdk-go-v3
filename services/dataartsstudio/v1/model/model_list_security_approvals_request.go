package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityApprovalsRequest Request Object
type ListSecurityApprovalsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 申请人名称
	ProposerName *string `json:"proposer_name,omitempty"`

	// 工单id
	ApprovalId *string `json:"approval_id,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 工单状态
	StatusList *[]int32 `json:"status_list,omitempty"`

	// 申请开始时间
	ApplicationStartTime *int64 `json:"application_start_time,omitempty"`

	// 申请结束时间
	ApplicationEndTime *int64 `json:"application_end_time,omitempty"`

	// 升降序
	OrderByDesc *bool `json:"order_by_desc,omitempty"`

	// 排序参数, START_TIME,END_TIME
	OrderBy *ListSecurityApprovalsRequestOrderBy `json:"order_by,omitempty"`
}

func (o ListSecurityApprovalsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityApprovalsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityApprovalsRequest", string(data)}, " ")
}

type ListSecurityApprovalsRequestOrderBy struct {
	value string
}

type ListSecurityApprovalsRequestOrderByEnum struct {
	START_TIME ListSecurityApprovalsRequestOrderBy
	END_TIME   ListSecurityApprovalsRequestOrderBy
}

func GetListSecurityApprovalsRequestOrderByEnum() ListSecurityApprovalsRequestOrderByEnum {
	return ListSecurityApprovalsRequestOrderByEnum{
		START_TIME: ListSecurityApprovalsRequestOrderBy{
			value: "START_TIME",
		},
		END_TIME: ListSecurityApprovalsRequestOrderBy{
			value: "END_TIME",
		},
	}
}

func (c ListSecurityApprovalsRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityApprovalsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityApprovalsRequestOrderBy) UnmarshalJSON(b []byte) error {
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

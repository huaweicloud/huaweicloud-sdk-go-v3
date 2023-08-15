package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InvitationDetail 被邀请列表
type InvitationDetail struct {

	// 被邀请方租户名，IAM用户名
	InvitedUser string `json:"invited_user"`

	// 邀请状态，可选：已退出（quit），等待中（waiting），已拒绝（reject），已解散（released），其他状态不允许删除
	Status *InvitationDetailStatus `json:"status,omitempty"`

	// 被邀请方服务实例ID
	InvitedBcsId *string `json:"invited_bcs_id,omitempty"`
}

func (o InvitationDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvitationDetail struct{}"
	}

	return strings.Join([]string{"InvitationDetail", string(data)}, " ")
}

type InvitationDetailStatus struct {
	value string
}

type InvitationDetailStatusEnum struct {
	QUIT     InvitationDetailStatus
	WAITING  InvitationDetailStatus
	REJECT   InvitationDetailStatus
	RELEASED InvitationDetailStatus
}

func GetInvitationDetailStatusEnum() InvitationDetailStatusEnum {
	return InvitationDetailStatusEnum{
		QUIT: InvitationDetailStatus{
			value: "quit",
		},
		WAITING: InvitationDetailStatus{
			value: "waiting",
		},
		REJECT: InvitationDetailStatus{
			value: "reject",
		},
		RELEASED: InvitationDetailStatus{
			value: "released",
		},
	}
}

func (c InvitationDetailStatus) Value() string {
	return c.value
}

func (c InvitationDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InvitationDetailStatus) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 被邀请列表
type InvitationDetail struct {
	// 被邀请方租户名，IAM用户名

	InvitedUser string `json:"invited_user"`
	// 邀请状态，可选：已退出（quit），等待中（waiting），已拒绝（reject），已解散（released），其他状态不允许删除

	Status *InvitationDetailStatus `json:"status,omitempty"`
	// 被邀请方bcs实例id

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

func (c InvitationDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InvitationDetailStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Member
type Member struct {

	// 共享状态
	Status MemberStatus `json:"status"`

	// 共享时间，例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间，例如:\"2020-02-05T10:38:34.209782\"
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 备份副本id
	BackupId *string `json:"backup_id,omitempty"`

	// 接受的共享备份副本注册的镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 接受备份共享的项目id
	DestProjectId *string `json:"dest_project_id,omitempty"`

	// 目标端接受共享备份的存储库id
	VaultId *string `json:"vault_id,omitempty"`

	// 共享记录id
	Id *string `json:"id,omitempty"`
}

func (o Member) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}

type MemberStatus struct {
	value string
}

type MemberStatusEnum struct {
	PENDING  MemberStatus
	ACCEPTED MemberStatus
	REJECTED MemberStatus
}

func GetMemberStatusEnum() MemberStatusEnum {
	return MemberStatusEnum{
		PENDING: MemberStatus{
			value: "pending",
		},
		ACCEPTED: MemberStatus{
			value: "accepted",
		},
		REJECTED: MemberStatus{
			value: "rejected",
		},
	}
}

func (c MemberStatus) Value() string {
	return c.value
}

func (c MemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberStatus) UnmarshalJSON(b []byte) error {
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

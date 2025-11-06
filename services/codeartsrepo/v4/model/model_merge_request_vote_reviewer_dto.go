package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MergeRequestVoteReviewerDto 合并请求打分模式评审人
type MergeRequestVoteReviewerDto struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户名。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 用户状态。 **取值范围：** - active: 可用账户。 - blocked: 被锁定用户。 - error: 未查询到该用户。
	State *MergeRequestVoteReviewerDtoState `json:"state,omitempty"`

	// 服务级权限状态 0：停用 1：启用
	ServiceLicenseStatus *int32 `json:"service_license_status,omitempty"`

	// 用户头像url
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// 用户头像路径
	AvatarPath *string `json:"avatar_path,omitempty"`

	// 用户邮箱
	Email *string `json:"email,omitempty"`

	// 用户中文名称
	NameCn *string `json:"name_cn,omitempty"`

	// 用户个人首页
	WebUrl *string `json:"web_url,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 部分查询接口校验到传参里的用户权限不足或不存在时，返回该用户但该字段不为空用于提示。
	ErrorMessage *string `json:"error_message,omitempty"`

	// **参数解释：** 是否为committer。
	IsCommitter *bool `json:"is_committer,omitempty"`
}

func (o MergeRequestVoteReviewerDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestVoteReviewerDto struct{}"
	}

	return strings.Join([]string{"MergeRequestVoteReviewerDto", string(data)}, " ")
}

type MergeRequestVoteReviewerDtoState struct {
	value string
}

type MergeRequestVoteReviewerDtoStateEnum struct {
	ACTIVE  MergeRequestVoteReviewerDtoState
	BLOCKED MergeRequestVoteReviewerDtoState
	ERROR   MergeRequestVoteReviewerDtoState
}

func GetMergeRequestVoteReviewerDtoStateEnum() MergeRequestVoteReviewerDtoStateEnum {
	return MergeRequestVoteReviewerDtoStateEnum{
		ACTIVE: MergeRequestVoteReviewerDtoState{
			value: "active",
		},
		BLOCKED: MergeRequestVoteReviewerDtoState{
			value: "blocked",
		},
		ERROR: MergeRequestVoteReviewerDtoState{
			value: "error",
		},
	}
}

func (c MergeRequestVoteReviewerDtoState) Value() string {
	return c.value
}

func (c MergeRequestVoteReviewerDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MergeRequestVoteReviewerDtoState) UnmarshalJSON(b []byte) error {
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

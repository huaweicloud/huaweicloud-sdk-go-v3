package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobUpdateRecordListVoResultJobUpdateRecords struct {

	// 修改编号
	Id *string `json:"id,omitempty"`

	// 更新编号
	UpdateNumber *int32 `json:"update_number,omitempty"`

	// 类型
	UpdateType *string `json:"update_type,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`

	// 更新时间
	UpdateAt *int64 `json:"update_at,omitempty"`
}

func (o JobUpdateRecordListVoResultJobUpdateRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobUpdateRecordListVoResultJobUpdateRecords struct{}"
	}

	return strings.Join([]string{"JobUpdateRecordListVoResultJobUpdateRecords", string(data)}, " ")
}

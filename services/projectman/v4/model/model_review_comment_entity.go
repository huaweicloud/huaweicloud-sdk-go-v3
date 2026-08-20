package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReviewCommentEntity struct {

	// 评审意见对象ID。
	Id *string `json:"id,omitempty"`

	// 评审用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 其他用户Id（转他人）。
	OtherUserId *string `json:"other_user_id,omitempty"`

	// 评审结果。
	Result *string `json:"result,omitempty"`

	// 评审意见。
	Comment *string `json:"comment,omitempty"`
}

func (o ReviewCommentEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewCommentEntity struct{}"
	}

	return strings.Join([]string{"ReviewCommentEntity", string(data)}, " ")
}

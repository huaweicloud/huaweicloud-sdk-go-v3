package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalReviewerInfoDto struct {

	// **参数解释：** 需要更新的检视人ID列表 数字以英文逗号分隔 111,222 **取值范围：** 不涉及。
	ReviewerIds *string `json:"reviewer_ids,omitempty"`
}

func (o ApprovalReviewerInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalReviewerInfoDto struct{}"
	}

	return strings.Join([]string{"ApprovalReviewerInfoDto", string(data)}, " ")
}

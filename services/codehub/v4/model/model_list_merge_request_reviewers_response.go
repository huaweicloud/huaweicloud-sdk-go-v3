package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestReviewersResponse Response Object
type ListMergeRequestReviewersResponse struct {

	// **参数解释：** 必选检视人列表。 **取值范围：** 不涉及。
	RequiredReviewersList *[]UserBasicDto `json:"required_reviewers_list,omitempty"`

	// **参数解释：** 可选检视人列表。 **取值范围：** 不涉及。
	OptionalReviewersList *[]UserBasicDto `json:"optional_reviewers_list,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMergeRequestReviewersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestReviewersResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestReviewersResponse", string(data)}, " ")
}

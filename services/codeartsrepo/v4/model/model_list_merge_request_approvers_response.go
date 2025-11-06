package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestApproversResponse Response Object
type ListMergeRequestApproversResponse struct {

	// **参数解释：** 必选审核人列表。 **取值范围：** 不涉及。
	RequiredApproversList *[]UserBasicDto `json:"required_approvers_list,omitempty"`

	// **参数解释：** 可选审核人列表。 **取值范围：** 不涉及。
	OptionalApproversList *[]UserBasicDto `json:"optional_approvers_list,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMergeRequestApproversResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestApproversResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestApproversResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMergeRequestApproversResponse Response Object
type UpdateMergeRequestApproversResponse struct {
	Body           *[]ApproverBasicDto `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateMergeRequestApproversResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMergeRequestApproversResponse struct{}"
	}

	return strings.Join([]string{"UpdateMergeRequestApproversResponse", string(data)}, " ")
}

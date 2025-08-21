package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResolveMergeRequestConflictsResponse Response Object
type ResolveMergeRequestConflictsResponse struct {

	// **参数解释：** prompt property name already exists
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResolveMergeRequestConflictsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResolveMergeRequestConflictsResponse struct{}"
	}

	return strings.Join([]string{"ResolveMergeRequestConflictsResponse", string(data)}, " ")
}

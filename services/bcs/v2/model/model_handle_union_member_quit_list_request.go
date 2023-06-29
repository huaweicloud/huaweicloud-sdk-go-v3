package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleUnionMemberQuitListRequest Request Object
type HandleUnionMemberQuitListRequest struct {
	Body *QuitUnionFromMemberListRequestBody `json:"body,omitempty"`
}

func (o HandleUnionMemberQuitListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleUnionMemberQuitListRequest struct{}"
	}

	return strings.Join([]string{"HandleUnionMemberQuitListRequest", string(data)}, " ")
}

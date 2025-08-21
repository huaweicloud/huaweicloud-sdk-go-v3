package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMergeRequestVoteResponse Response Object
type DeleteMergeRequestVoteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMergeRequestVoteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMergeRequestVoteResponse struct{}"
	}

	return strings.Join([]string{"DeleteMergeRequestVoteResponse", string(data)}, " ")
}

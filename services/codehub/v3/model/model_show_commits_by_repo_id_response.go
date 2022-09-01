package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCommitsByRepoIdResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *CommitList `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCommitsByRepoIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitsByRepoIdResponse struct{}"
	}

	return strings.Join([]string{"ShowCommitsByRepoIdResponse", string(data)}, " ")
}

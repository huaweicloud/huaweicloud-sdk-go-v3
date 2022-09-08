package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSingleCommitResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *SpecificCommitInfo `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSingleCommitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleCommitResponse struct{}"
	}

	return strings.Join([]string{"ShowSingleCommitResponse", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDiffCommitResponse struct {
	Error *Error `json:"error,omitempty"`
	// 差异列表

	Result *[]DiffCommitInfo `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDiffCommitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiffCommitResponse struct{}"
	}

	return strings.Join([]string{"ShowDiffCommitResponse", string(data)}, " ")
}

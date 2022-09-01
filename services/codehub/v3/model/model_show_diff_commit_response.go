package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDiffCommitResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	// 差异列表
	Result *[]DiffCommitInfo `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDiffCommitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiffCommitResponse struct{}"
	}

	return strings.Join([]string{"ShowDiffCommitResponse", string(data)}, " ")
}

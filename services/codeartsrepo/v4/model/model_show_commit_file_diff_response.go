package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommitFileDiffResponse Response Object
type ShowCommitFileDiffResponse struct {

	// 文件差异信息列表
	Body           *[]FileDiffDto `json:"body,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowCommitFileDiffResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitFileDiffResponse struct{}"
	}

	return strings.Join([]string{"ShowCommitFileDiffResponse", string(data)}, " ")
}

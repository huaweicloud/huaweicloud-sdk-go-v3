package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadIpdImageInIssueResponse Response Object
type UploadIpdImageInIssueResponse struct {

	// 请求状态,成功为success，失败时会抛出异常
	Status *string `json:"status,omitempty"`

	// 请求信息，一般情况为空。
	Message *string `json:"message,omitempty"`

	Result         *IssueEntity `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UploadIpdImageInIssueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadIpdImageInIssueResponse struct{}"
	}

	return strings.Join([]string{"UploadIpdImageInIssueResponse", string(data)}, " ")
}

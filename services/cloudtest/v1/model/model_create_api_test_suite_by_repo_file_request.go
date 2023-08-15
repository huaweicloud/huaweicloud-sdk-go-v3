package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApiTestSuiteByRepoFileRequest Request Object
type CreateApiTestSuiteByRepoFileRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CreateTestSuitByRepoFileInfo `json:"body,omitempty"`
}

func (o CreateApiTestSuiteByRepoFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiTestSuiteByRepoFileRequest struct{}"
	}

	return strings.Join([]string{"CreateApiTestSuiteByRepoFileRequest", string(data)}, " ")
}

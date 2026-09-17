package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTestVersionCaseRequest Request Object
type CreateTestVersionCaseRequest struct {

	// 分支或者迭代uri
	VersionUri string `json:"version_uri"`

	Body *TestCaseInfo `json:"body,omitempty"`
}

func (o CreateTestVersionCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestVersionCaseRequest struct{}"
	}

	return strings.Join([]string{"CreateTestVersionCaseRequest", string(data)}, " ")
}

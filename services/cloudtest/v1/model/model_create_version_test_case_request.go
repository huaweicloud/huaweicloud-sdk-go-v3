package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVersionTestCaseRequest Request Object
type CreateVersionTestCaseRequest struct {

	// 分支或者迭代uri
	VersionId string `json:"version_id"`

	Body *ApiResultTestCaseVo `json:"body,omitempty"`
}

func (o CreateVersionTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVersionTestCaseRequest struct{}"
	}

	return strings.Join([]string{"CreateVersionTestCaseRequest", string(data)}, " ")
}

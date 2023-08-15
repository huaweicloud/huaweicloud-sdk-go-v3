package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTestCaseRequest Request Object
type BatchDeleteTestCaseRequest struct {

	// 项目唯一标识，固定长度32位字符
	ProjectId string `json:"project_id"`

	Body *BatchDeleteTestCaseRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTestCaseRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestCaseRequest", string(data)}, " ")
}

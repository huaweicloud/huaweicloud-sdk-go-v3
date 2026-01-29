package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteIssuesV4Request Request Object
type BatchDeleteIssuesV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *BatchDeleteIssuesRequestV4 `json:"body,omitempty"`
}

func (o BatchDeleteIssuesV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIssuesV4Request struct{}"
	}

	return strings.Join([]string{"BatchDeleteIssuesV4Request", string(data)}, " ")
}

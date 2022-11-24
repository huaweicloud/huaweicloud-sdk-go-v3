package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteIssuesV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteIssuesV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIssuesV4Response struct{}"
	}

	return strings.Join([]string{"BatchDeleteIssuesV4Response", string(data)}, " ")
}

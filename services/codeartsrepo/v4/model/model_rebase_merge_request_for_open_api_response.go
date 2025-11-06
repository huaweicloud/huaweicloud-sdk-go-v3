package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebaseMergeRequestForOpenApiResponse Response Object
type RebaseMergeRequestForOpenApiResponse struct {

	// 响应信息
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RebaseMergeRequestForOpenApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebaseMergeRequestForOpenApiResponse struct{}"
	}

	return strings.Join([]string{"RebaseMergeRequestForOpenApiResponse", string(data)}, " ")
}

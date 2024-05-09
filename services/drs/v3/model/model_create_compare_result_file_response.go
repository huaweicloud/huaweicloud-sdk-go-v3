package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCompareResultFileResponse Response Object
type CreateCompareResultFileResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateCompareResultFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareResultFileResponse struct{}"
	}

	return strings.Join([]string{"CreateCompareResultFileResponse", string(data)}, " ")
}

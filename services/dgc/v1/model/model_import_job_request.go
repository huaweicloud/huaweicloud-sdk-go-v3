package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportJobRequest struct {
	Body *ImportFileReq `json:"body,omitempty"`
}

func (o ImportJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportJobRequest struct{}"
	}

	return strings.Join([]string{"ImportJobRequest", string(data)}, " ")
}

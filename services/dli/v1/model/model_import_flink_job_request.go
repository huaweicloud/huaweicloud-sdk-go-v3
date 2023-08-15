package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFlinkJobRequest Request Object
type ImportFlinkJobRequest struct {
	Body *ImportFlinkJobRequestBody `json:"body,omitempty"`
}

func (o ImportFlinkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFlinkJobRequest struct{}"
	}

	return strings.Join([]string{"ImportFlinkJobRequest", string(data)}, " ")
}

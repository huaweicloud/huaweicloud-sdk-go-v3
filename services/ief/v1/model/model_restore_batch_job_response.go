package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RestoreBatchJobResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreBatchJobResponse struct{}"
	}

	return strings.Join([]string{"RestoreBatchJobResponse", string(data)}, " ")
}

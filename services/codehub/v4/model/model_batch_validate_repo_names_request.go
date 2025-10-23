package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchValidateRepoNamesRequest Request Object
type BatchValidateRepoNamesRequest struct {
	Body *ValidateRepoNameDto `json:"body,omitempty"`
}

func (o BatchValidateRepoNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchValidateRepoNamesRequest struct{}"
	}

	return strings.Join([]string{"BatchValidateRepoNamesRequest", string(data)}, " ")
}

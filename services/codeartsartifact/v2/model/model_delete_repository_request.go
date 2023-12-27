package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepositoryRequest Request Object
type DeleteRepositoryRequest struct {
	Body *IdeRepoRevisionModel `json:"body,omitempty"`
}

func (o DeleteRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepositoryRequest struct{}"
	}

	return strings.Join([]string{"DeleteRepositoryRequest", string(data)}, " ")
}

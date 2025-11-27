package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepoRequest Request Object
type UpdateRepoRequest struct {

	// 仓库id
	Repoid string `json:"repoid"`

	Body *UpdateRepoRequestBody `json:"body,omitempty"`
}

func (o UpdateRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoRequest struct{}"
	}

	return strings.Join([]string{"UpdateRepoRequest", string(data)}, " ")
}

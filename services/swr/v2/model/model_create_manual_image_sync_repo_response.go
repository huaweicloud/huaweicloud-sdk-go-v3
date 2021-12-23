package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateManualImageSyncRepoResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateManualImageSyncRepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualImageSyncRepoResponse struct{}"
	}

	return strings.Join([]string{"CreateManualImageSyncRepoResponse", string(data)}, " ")
}

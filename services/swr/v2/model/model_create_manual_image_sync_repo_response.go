package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManualImageSyncRepoResponse Response Object
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

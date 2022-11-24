package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateImageSyncRepoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateImageSyncRepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageSyncRepoResponse struct{}"
	}

	return strings.Join([]string{"CreateImageSyncRepoResponse", string(data)}, " ")
}

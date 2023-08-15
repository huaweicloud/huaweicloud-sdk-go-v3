package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageSyncRepoResponse Response Object
type DeleteImageSyncRepoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteImageSyncRepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageSyncRepoResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageSyncRepoResponse", string(data)}, " ")
}

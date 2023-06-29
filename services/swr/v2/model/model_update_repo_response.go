package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepoResponse Response Object
type UpdateRepoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepoResponse", string(data)}, " ")
}

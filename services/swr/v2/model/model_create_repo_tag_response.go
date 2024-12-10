package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRepoTagResponse Response Object
type CreateRepoTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRepoTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoTagResponse struct{}"
	}

	return strings.Join([]string{"CreateRepoTagResponse", string(data)}, " ")
}

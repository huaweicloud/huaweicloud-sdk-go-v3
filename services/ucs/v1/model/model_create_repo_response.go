package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRepoResponse Response Object
type CreateRepoResponse struct {

	// 仓库id
	Uid            *string `json:"uid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoResponse struct{}"
	}

	return strings.Join([]string{"CreateRepoResponse", string(data)}, " ")
}

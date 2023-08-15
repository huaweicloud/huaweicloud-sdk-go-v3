package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepoResponse Response Object
type DeleteRepoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepoResponse struct{}"
	}

	return strings.Join([]string{"DeleteRepoResponse", string(data)}, " ")
}

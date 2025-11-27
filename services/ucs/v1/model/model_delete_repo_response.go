package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepoResponse Response Object
type DeleteRepoResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepoResponse struct{}"
	}

	return strings.Join([]string{"DeleteRepoResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepoTagResponse Response Object
type DeleteRepoTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRepoTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepoTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteRepoTagResponse", string(data)}, " ")
}

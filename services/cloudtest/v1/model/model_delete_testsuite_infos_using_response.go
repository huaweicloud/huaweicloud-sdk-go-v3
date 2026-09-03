package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTestsuiteInfosUsingResponse Response Object
type DeleteTestsuiteInfosUsingResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTestsuiteInfosUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTestsuiteInfosUsingResponse struct{}"
	}

	return strings.Join([]string{"DeleteTestsuiteInfosUsingResponse", string(data)}, " ")
}

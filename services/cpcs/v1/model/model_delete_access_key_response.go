package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccessKeyResponse Response Object
type DeleteAccessKeyResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteAccessKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAccessKeyResponse", string(data)}, " ")
}

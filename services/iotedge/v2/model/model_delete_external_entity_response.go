package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExternalEntityResponse Response Object
type DeleteExternalEntityResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteExternalEntityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExternalEntityResponse struct{}"
	}

	return strings.Join([]string{"DeleteExternalEntityResponse", string(data)}, " ")
}

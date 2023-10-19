package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubAppResponse Response Object
type DeleteSubAppResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSubAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubAppResponse", string(data)}, " ")
}

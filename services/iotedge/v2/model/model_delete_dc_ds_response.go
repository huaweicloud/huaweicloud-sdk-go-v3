package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDcDsResponse Response Object
type DeleteDcDsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDcDsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDcDsResponse struct{}"
	}

	return strings.Join([]string{"DeleteDcDsResponse", string(data)}, " ")
}

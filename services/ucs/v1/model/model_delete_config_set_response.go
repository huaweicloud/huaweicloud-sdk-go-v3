package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConfigSetResponse Response Object
type DeleteConfigSetResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteConfigSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteConfigSetResponse", string(data)}, " ")
}

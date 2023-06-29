package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRsuResponse Response Object
type DeleteRsuResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRsuResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRsuResponse struct{}"
	}

	return strings.Join([]string{"DeleteRsuResponse", string(data)}, " ")
}

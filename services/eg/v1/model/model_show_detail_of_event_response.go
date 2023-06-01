package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDetailOfEventResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailOfEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventResponse", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRackResponse Response Object
type ShowRackResponse struct {
	Rack           *Rack `json:"rack,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowRackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRackResponse struct{}"
	}

	return strings.Join([]string{"ShowRackResponse", string(data)}, " ")
}

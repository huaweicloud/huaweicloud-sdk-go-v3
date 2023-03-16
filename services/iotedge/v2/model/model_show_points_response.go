package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPointsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPointsResponse struct{}"
	}

	return strings.Join([]string{"ShowPointsResponse", string(data)}, " ")
}

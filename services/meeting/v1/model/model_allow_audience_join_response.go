package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowAudienceJoinResponse Response Object
type AllowAudienceJoinResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AllowAudienceJoinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowAudienceJoinResponse struct{}"
	}

	return strings.Join([]string{"AllowAudienceJoinResponse", string(data)}, " ")
}

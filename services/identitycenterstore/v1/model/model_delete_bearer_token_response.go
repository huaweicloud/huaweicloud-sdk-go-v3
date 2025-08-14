package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBearerTokenResponse Response Object
type DeleteBearerTokenResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBearerTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBearerTokenResponse struct{}"
	}

	return strings.Join([]string{"DeleteBearerTokenResponse", string(data)}, " ")
}

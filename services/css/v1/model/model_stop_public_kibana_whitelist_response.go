package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopPublicKibanaWhitelistResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopPublicKibanaWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPublicKibanaWhitelistResponse struct{}"
	}

	return strings.Join([]string{"StopPublicKibanaWhitelistResponse", string(data)}, " ")
}

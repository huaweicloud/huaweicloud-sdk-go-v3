package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetSiteAuthConfigResponse Response Object
type ResetSiteAuthConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetSiteAuthConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetSiteAuthConfigResponse struct{}"
	}

	return strings.Join([]string{"ResetSiteAuthConfigResponse", string(data)}, " ")
}

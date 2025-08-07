package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEpConfigsResponse Response Object
type ShowEpConfigsResponse struct {
	SupportItem    *EnterpriseProjectConfig `json:"support_item,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowEpConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEpConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowEpConfigsResponse", string(data)}, " ")
}

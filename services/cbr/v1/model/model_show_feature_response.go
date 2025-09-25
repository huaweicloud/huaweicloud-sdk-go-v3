package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFeatureResponse Response Object
type ShowFeatureResponse struct {
	Body           map[string]interface{} `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowFeatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFeatureResponse struct{}"
	}

	return strings.Join([]string{"ShowFeatureResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeaturesResponse Response Object
type ListFeaturesResponse struct {
	Body           map[string]interface{} `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListFeaturesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeaturesResponse struct{}"
	}

	return strings.Join([]string{"ListFeaturesResponse", string(data)}, " ")
}

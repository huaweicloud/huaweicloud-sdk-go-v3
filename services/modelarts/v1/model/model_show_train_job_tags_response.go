package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainJobTagsResponse Response Object
type ShowTrainJobTagsResponse struct {
	Tags           *[]TmsTag `json:"tags,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTrainJobTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainJobTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainJobTagsResponse", string(data)}, " ")
}

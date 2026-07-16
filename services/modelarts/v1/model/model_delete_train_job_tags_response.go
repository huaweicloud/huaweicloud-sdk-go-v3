package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrainJobTagsResponse Response Object
type DeleteTrainJobTagsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTrainJobTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrainJobTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrainJobTagsResponse", string(data)}, " ")
}

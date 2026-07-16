package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainJobTagsResponse Response Object
type CreateTrainJobTagsResponse struct {

	// **参数解释**：TMS标签列表。
	Tags           *[]TmsTagResp `json:"tags,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateTrainJobTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainJobTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateTrainJobTagsResponse", string(data)}, " ")
}

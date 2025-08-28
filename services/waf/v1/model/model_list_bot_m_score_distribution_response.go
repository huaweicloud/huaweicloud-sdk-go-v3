package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBotMScoreDistributionResponse Response Object
type ListBotMScoreDistributionResponse struct {
	Body           *[]TypedStatBucket `json:"body,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListBotMScoreDistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBotMScoreDistributionResponse struct{}"
	}

	return strings.Join([]string{"ListBotMScoreDistributionResponse", string(data)}, " ")
}

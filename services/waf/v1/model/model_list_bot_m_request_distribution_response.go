package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBotMRequestDistributionResponse Response Object
type ListBotMRequestDistributionResponse struct {
	NormalBucket *BotRequestDistributionsNormalBucket `json:"normal_bucket,omitempty"`

	BotClassificationBucket *BotTypeDistributions `json:"bot_classification_bucket,omitempty"`

	ActionBucket   *ActionDistributions `json:"action_bucket,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListBotMRequestDistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBotMRequestDistributionResponse struct{}"
	}

	return strings.Join([]string{"ListBotMRequestDistributionResponse", string(data)}, " ")
}

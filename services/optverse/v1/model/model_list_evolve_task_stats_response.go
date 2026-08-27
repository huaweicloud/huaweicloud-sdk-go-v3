package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEvolveTaskStatsResponse Response Object
type ListEvolveTaskStatsResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEvolveTaskStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEvolveTaskStatsResponse struct{}"
	}

	return strings.Join([]string{"ListEvolveTaskStatsResponse", string(data)}, " ")
}

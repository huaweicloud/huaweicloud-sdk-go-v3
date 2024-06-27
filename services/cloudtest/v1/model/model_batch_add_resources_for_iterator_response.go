package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddResourcesForIteratorResponse Response Object
type BatchAddResourcesForIteratorResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueStringForOk `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchAddResourcesForIteratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddResourcesForIteratorResponse struct{}"
	}

	return strings.Join([]string{"BatchAddResourcesForIteratorResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceSetsResponse Response Object
type ListServiceSetsResponse struct {
	Data           *ServiceSetRecords `json:"data,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListServiceSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSetsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceSetsResponse", string(data)}, " ")
}

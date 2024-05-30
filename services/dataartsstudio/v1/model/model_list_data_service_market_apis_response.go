package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataServiceMarketApisResponse Response Object
type ListDataServiceMarketApisResponse struct {

	// API总条数。
	Total *int32 `json:"total,omitempty"`

	// API列表。
	Apis           *[]AdvancedMallApiDto `json:"apis,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListDataServiceMarketApisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataServiceMarketApisResponse struct{}"
	}

	return strings.Join([]string{"ListDataServiceMarketApisResponse", string(data)}, " ")
}

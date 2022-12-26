package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHistorySqlResponse struct {
	Results        *[]QuertHistorySqlResultsBody `json:"results,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListHistorySqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistorySqlResponse struct{}"
	}

	return strings.Join([]string{"ListHistorySqlResponse", string(data)}, " ")
}

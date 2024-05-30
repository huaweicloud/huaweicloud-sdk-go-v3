package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountOverviewsResponse Response Object
type CountOverviewsResponse struct {
	Data           *CountOverviewsResultData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o CountOverviewsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountOverviewsResponse struct{}"
	}

	return strings.Join([]string{"CountOverviewsResponse", string(data)}, " ")
}

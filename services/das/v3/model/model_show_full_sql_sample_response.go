package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFullSqlSampleResponse Response Object
type ShowFullSqlSampleResponse struct {
	Sample         *FullSqlSampleInfo `json:"sample,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowFullSqlSampleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFullSqlSampleResponse struct{}"
	}

	return strings.Join([]string{"ShowFullSqlSampleResponse", string(data)}, " ")
}

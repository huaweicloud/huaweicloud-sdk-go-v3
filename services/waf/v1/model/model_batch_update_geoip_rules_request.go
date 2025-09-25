package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateGeoipRulesRequest Request Object
type BatchUpdateGeoipRulesRequest struct {
	Body *BatchUpdateGeoipRulesRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateGeoipRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateGeoipRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateGeoipRulesRequest", string(data)}, " ")
}

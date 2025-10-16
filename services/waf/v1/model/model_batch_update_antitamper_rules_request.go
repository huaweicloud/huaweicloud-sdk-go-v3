package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAntitamperRulesRequest Request Object
type BatchUpdateAntitamperRulesRequest struct {
	Body *BatchUpdateAntiTamperRulesRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateAntitamperRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAntitamperRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateAntitamperRulesRequest", string(data)}, " ")
}

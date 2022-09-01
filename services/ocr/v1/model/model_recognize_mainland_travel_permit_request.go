package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeMainlandTravelPermitRequest struct {
	Body *MainlandTravelPermitRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeMainlandTravelPermitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMainlandTravelPermitRequest struct{}"
	}

	return strings.Join([]string{"RecognizeMainlandTravelPermitRequest", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeMainlandTravelPermitResponse Response Object
type RecognizeMainlandTravelPermitResponse struct {
	Result         *MainlandTravelPermitResult `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o RecognizeMainlandTravelPermitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMainlandTravelPermitResponse struct{}"
	}

	return strings.Join([]string{"RecognizeMainlandTravelPermitResponse", string(data)}, " ")
}

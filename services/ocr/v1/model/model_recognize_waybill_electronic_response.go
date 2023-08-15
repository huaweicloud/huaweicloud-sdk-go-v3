package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeWaybillElectronicResponse Response Object
type RecognizeWaybillElectronicResponse struct {
	Result         *WaybillElectronicResult `json:"result,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o RecognizeWaybillElectronicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeWaybillElectronicResponse struct{}"
	}

	return strings.Join([]string{"RecognizeWaybillElectronicResponse", string(data)}, " ")
}

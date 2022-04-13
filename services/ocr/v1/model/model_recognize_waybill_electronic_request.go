package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeWaybillElectronicRequest struct {
	Body *WaybillElectronicRequestBody `json:"body,omitempty"`
}

func (o RecognizeWaybillElectronicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeWaybillElectronicRequest struct{}"
	}

	return strings.Join([]string{"RecognizeWaybillElectronicRequest", string(data)}, " ")
}

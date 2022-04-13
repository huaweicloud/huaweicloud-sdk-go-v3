package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeGeneralTableRequest struct {
	Body *GeneralTableRequestBody `json:"body,omitempty"`
}

func (o RecognizeGeneralTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTableRequest struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTableRequest", string(data)}, " ")
}

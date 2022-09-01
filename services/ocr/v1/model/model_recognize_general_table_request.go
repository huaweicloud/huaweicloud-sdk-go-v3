package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeGeneralTableRequest struct {
	Body *GeneralTableRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeGeneralTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTableRequest struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTableRequest", string(data)}, " ")
}

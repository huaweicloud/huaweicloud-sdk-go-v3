package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeGeneralTableResponse struct {
	Result         *GeneralTableResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o RecognizeGeneralTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTableResponse struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTableResponse", string(data)}, " ")
}

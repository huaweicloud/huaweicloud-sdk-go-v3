package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeBankcardResponse struct {
	Result         *BankcardResult `json:"result,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RecognizeBankcardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeBankcardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeBankcardResponse", string(data)}, " ")
}

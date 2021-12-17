package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeMyanmarIdcardResponse struct {
	Result         *MyanmarIdcardResult `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o RecognizeMyanmarIdcardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMyanmarIdcardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeMyanmarIdcardResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeMyanmarIdcardResponse Response Object
type RecognizeMyanmarIdcardResponse struct {
	Result *MyanmarIdcardResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeMyanmarIdcardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMyanmarIdcardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeMyanmarIdcardResponse", string(data)}, " ")
}

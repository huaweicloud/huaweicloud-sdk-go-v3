package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeMyanmarIdcardRequest struct {
	Body *MyanmarIdcardRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeMyanmarIdcardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMyanmarIdcardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeMyanmarIdcardRequest", string(data)}, " ")
}

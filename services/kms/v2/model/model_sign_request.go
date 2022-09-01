package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SignRequest struct {
	Body *SignRequestBody `json:"body,omitempty" xml:"body"`
}

func (o SignRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignRequest struct{}"
	}

	return strings.Join([]string{"SignRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomAnswersRequest Request Object
type CreateCustomAnswersRequest struct {
	Body *AiguardCustomAnswersBody `json:"body,omitempty"`
}

func (o CreateCustomAnswersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomAnswersRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomAnswersRequest", string(data)}, " ")
}

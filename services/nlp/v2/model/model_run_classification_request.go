package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunClassificationRequest Request Object
type RunClassificationRequest struct {
	Body *ClassificationReq `json:"body,omitempty"`
}

func (o RunClassificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunClassificationRequest struct{}"
	}

	return strings.Join([]string{"RunClassificationRequest", string(data)}, " ")
}

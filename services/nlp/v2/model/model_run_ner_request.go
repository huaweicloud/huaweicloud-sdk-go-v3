package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunNerRequest struct {
	Body *NerRequest `json:"body,omitempty" xml:"body"`
}

func (o RunNerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunNerRequest struct{}"
	}

	return strings.Join([]string{"RunNerRequest", string(data)}, " ")
}

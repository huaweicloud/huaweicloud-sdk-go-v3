package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunDocClassificationRequest struct {
	Body *DocumentClassificationReq `json:"body,omitempty" xml:"body"`
}

func (o RunDocClassificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDocClassificationRequest struct{}"
	}

	return strings.Join([]string{"RunDocClassificationRequest", string(data)}, " ")
}

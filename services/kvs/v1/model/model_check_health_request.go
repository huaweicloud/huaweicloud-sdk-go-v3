package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckHealthRequest Request Object
type CheckHealthRequest struct {
	Body *CheckHealthRequestBody `bson:"body,omitempty"`
}

func (o CheckHealthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckHealthRequest struct{}"
	}

	return strings.Join([]string{"CheckHealthRequest", string(data)}, " ")
}

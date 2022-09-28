package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskInstancePodResponse struct {

	// pod信息
	Metadata       *string `json:"metadata,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskInstancePodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInstancePodResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskInstancePodResponse", string(data)}, " ")
}

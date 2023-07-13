package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBizMetricRequest Request Object
type DeleteBizMetricRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *IdsParam `json:"body,omitempty"`
}

func (o DeleteBizMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBizMetricRequest struct{}"
	}

	return strings.Join([]string{"DeleteBizMetricRequest", string(data)}, " ")
}

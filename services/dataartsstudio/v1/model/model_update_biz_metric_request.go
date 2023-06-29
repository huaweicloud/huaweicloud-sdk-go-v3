package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBizMetricRequest Request Object
type UpdateBizMetricRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BizMetricVo `json:"body,omitempty"`
}

func (o UpdateBizMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBizMetricRequest struct{}"
	}

	return strings.Join([]string{"UpdateBizMetricRequest", string(data)}, " ")
}

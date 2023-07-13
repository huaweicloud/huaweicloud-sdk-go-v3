package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBizMetricRequest Request Object
type CreateBizMetricRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BizMetricVo `json:"body,omitempty"`
}

func (o CreateBizMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBizMetricRequest struct{}"
	}

	return strings.Join([]string{"CreateBizMetricRequest", string(data)}, " ")
}

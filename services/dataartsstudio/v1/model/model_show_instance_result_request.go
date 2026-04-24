package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceResultRequest Request Object
type ShowInstanceResultRequest struct {

	// projectId
	InstanceId string `json:"instance_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 每页的记录数，取值范围为0~100。
	Limit *int64 `json:"limit,omitempty"`

	// 分页偏移量，最小值为0。
	Offset *int64 `json:"offset,omitempty"`
}

func (o ShowInstanceResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResultRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceResultRequest", string(data)}, " ")
}

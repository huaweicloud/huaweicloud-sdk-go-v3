package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentInstanceRegionId 中心网络附件对端实例的regionID。
type AttachmentInstanceRegionId struct {

	// 中心网络附件对端实例的regionID。
	AttachmentInstanceRegionId string `json:"attachment_instance_region_id"`
}

func (o AttachmentInstanceRegionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentInstanceRegionId struct{}"
	}

	return strings.Join([]string{"AttachmentInstanceRegionId", string(data)}, " ")
}

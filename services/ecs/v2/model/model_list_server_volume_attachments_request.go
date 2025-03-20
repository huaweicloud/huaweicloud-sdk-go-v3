package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerVolumeAttachmentsRequest Request Object
type ListServerVolumeAttachmentsRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ListServerVolumeAttachmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerVolumeAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"ListServerVolumeAttachmentsRequest", string(data)}, " ")
}

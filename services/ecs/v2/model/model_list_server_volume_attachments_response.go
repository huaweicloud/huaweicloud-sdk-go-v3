package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerVolumeAttachmentsResponse Response Object
type ListServerVolumeAttachmentsResponse struct {

	// 云服务器挂载信息列表
	VolumeAttachments *[]ServerVolumeAttachment `json:"volumeAttachments,omitempty"`
	HttpStatusCode    int                       `json:"-"`
}

func (o ListServerVolumeAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerVolumeAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"ListServerVolumeAttachmentsResponse", string(data)}, " ")
}

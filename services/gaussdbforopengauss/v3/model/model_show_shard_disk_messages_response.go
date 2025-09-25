package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShardDiskMessagesResponse Response Object
type ShowShardDiskMessagesResponse struct {

	// **参数解释**： 分片磁盘信息列表。
	GroupDiskInfos *[]GroupDiskInfoResult `json:"group_disk_infos,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowShardDiskMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShardDiskMessagesResponse struct{}"
	}

	return strings.Join([]string{"ShowShardDiskMessagesResponse", string(data)}, " ")
}

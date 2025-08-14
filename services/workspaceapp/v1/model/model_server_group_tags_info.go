package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerGroupTagsInfo 服务器组关联tags信息。
type ServerGroupTagsInfo struct {

	// 服务器组唯一标识。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 标签列表。
	Tags *[]TmsTag `json:"tags,omitempty"`
}

func (o ServerGroupTagsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerGroupTagsInfo struct{}"
	}

	return strings.Join([]string{"ServerGroupTagsInfo", string(data)}, " ")
}

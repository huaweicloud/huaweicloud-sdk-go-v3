package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHotspotSessionConfigResponse Response Object
type ListHotspotSessionConfigResponse struct {

	// 查询热点会话迁移配置。
	Items          *[]HotspotSessionMigrationConfig `json:"items,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListHotspotSessionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotspotSessionConfigResponse struct{}"
	}

	return strings.Join([]string{"ListHotspotSessionConfigResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReplaySupportSourceConfigResponse Response Object
type ListReplaySupportSourceConfigResponse struct {

	// 流量回放源数据配置信息
	ConfigList *[]ReplaySourceConfigResp `json:"config_list,omitempty"`

	// 总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListReplaySupportSourceConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplaySupportSourceConfigResponse struct{}"
	}

	return strings.Join([]string{"ListReplaySupportSourceConfigResponse", string(data)}, " ")
}

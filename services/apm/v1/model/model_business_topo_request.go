package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 获取应用拓扑数据的请求参数。
type BusinessTopoRequest struct {

	// 目标应用id。
	TargetBusinessId int64 `json:"target_business_id"`

	// 环境标签列表，可为空。
	EnvTagList *[]int64 `json:"env_tag_list,omitempty"`

	// 方向，可为空。
	Direction *string `json:"direction,omitempty"`

	// 结束时间。
	EndTime string `json:"end_time"`

	// 开始时间。
	StartTime string `json:"start_time"`

	// 过滤。
	FilterUser *bool `json:"filter_user,omitempty"`
}

func (o BusinessTopoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessTopoRequest struct{}"
	}

	return strings.Join([]string{"BusinessTopoRequest", string(data)}, " ")
}

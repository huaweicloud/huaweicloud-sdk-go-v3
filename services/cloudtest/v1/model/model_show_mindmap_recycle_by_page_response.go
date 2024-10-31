package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapRecycleByPageResponse Response Object
type ShowMindmapRecycleByPageResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Data *BasePageInfoMindmapRecycle `json:"data,omitempty"`

	// 错误信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMindmapRecycleByPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapRecycleByPageResponse struct{}"
	}

	return strings.Join([]string{"ShowMindmapRecycleByPageResponse", string(data)}, " ")
}

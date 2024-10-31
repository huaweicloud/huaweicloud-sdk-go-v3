package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapRecycleByPageRequest Request Object
type ShowMindmapRecycleByPageRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestMindmapRecyclePageParam `json:"body,omitempty"`
}

func (o ShowMindmapRecycleByPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapRecycleByPageRequest struct{}"
	}

	return strings.Join([]string{"ShowMindmapRecycleByPageRequest", string(data)}, " ")
}

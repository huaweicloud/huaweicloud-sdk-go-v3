package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapByPageRequest Request Object
type ShowMindmapByPageRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestMindmapPageParamV3 `json:"body,omitempty"`
}

func (o ShowMindmapByPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapByPageRequest struct{}"
	}

	return strings.Join([]string{"ShowMindmapByPageRequest", string(data)}, " ")
}

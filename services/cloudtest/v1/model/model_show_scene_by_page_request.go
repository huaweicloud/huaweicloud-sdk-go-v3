package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSceneByPageRequest Request Object
type ShowSceneByPageRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestScenePageParam `json:"body,omitempty"`
}

func (o ShowSceneByPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSceneByPageRequest struct{}"
	}

	return strings.Join([]string{"ShowSceneByPageRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTagRequest Request Object
type BatchTagRequest struct {

	// 实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Instance string `json:"instance"`

	Body *BatchRecommendationRequest `json:"body,omitempty"`
}

func (o BatchTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagRequest struct{}"
	}

	return strings.Join([]string{"BatchTagRequest", string(data)}, " ")
}

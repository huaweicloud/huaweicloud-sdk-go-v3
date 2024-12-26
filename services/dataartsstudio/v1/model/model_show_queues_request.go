package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueuesRequest Request Object
type ShowQueuesRequest struct {

	// 实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Instance string `json:"instance"`
}

func (o ShowQueuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueuesRequest struct{}"
	}

	return strings.Join([]string{"ShowQueuesRequest", string(data)}, " ")
}

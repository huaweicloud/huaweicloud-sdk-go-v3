package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSmnTopicInfoResponse Response Object
type ShowSmnTopicInfoResponse struct {

	// 主题信息
	Topics *[]SmnTopicInfo `json:"topics,omitempty"`

	// 主题总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSmnTopicInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmnTopicInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSmnTopicInfoResponse", string(data)}, " ")
}

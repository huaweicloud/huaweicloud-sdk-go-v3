package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPublishInfoResponse Response Object
type SearchPublishInfoResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 发布信息列表
	PublishMessages *[]ApiPublishDto `json:"publish_messages,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o SearchPublishInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPublishInfoResponse struct{}"
	}

	return strings.Join([]string{"SearchPublishInfoResponse", string(data)}, " ")
}

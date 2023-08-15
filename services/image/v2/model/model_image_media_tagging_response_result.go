package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageMediaTaggingResponseResult 调用成功时为图片标签内容。  调用失败时无此字段。
type ImageMediaTaggingResponseResult struct {

	// 标签列表集合。
	Tags *[]ImageMediaTaggingItemBody `json:"tags,omitempty"`
}

func (o ImageMediaTaggingResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMediaTaggingResponseResult struct{}"
	}

	return strings.Join([]string{"ImageMediaTaggingResponseResult", string(data)}, " ")
}

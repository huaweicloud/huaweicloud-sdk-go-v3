package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 调用成功时为图片标签内容。  调用失败时无此字段。
type ImageMediaTaggingDetResponseResult struct {

	// 标签列表集合。
	Tags *[]ImageMediaTaggingDetItemBody `json:"tags,omitempty"`
}

func (o ImageMediaTaggingDetResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMediaTaggingDetResponseResult struct{}"
	}

	return strings.Join([]string{"ImageMediaTaggingDetResponseResult", string(data)}, " ")
}

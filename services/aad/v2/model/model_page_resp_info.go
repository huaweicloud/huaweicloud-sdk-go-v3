package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageRespInfo struct {

	// 页面类型
	ContentType *string `json:"content_type,omitempty"`

	// 页面内容
	Content *string `json:"content,omitempty"`
}

func (o PageRespInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageRespInfo struct{}"
	}

	return strings.Join([]string{"PageRespInfo", string(data)}, " ")
}

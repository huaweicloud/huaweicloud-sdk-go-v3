package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageResp API列表查询响应  **⚠ 警告: 此Model不生成java代码，允许其它Model中引用，请勿直接作为Model使用**
type PageResp struct {

	// 总数
	Count *int32 `json:"count,omitempty"`
}

func (o PageResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageResp struct{}"
	}

	return strings.Join([]string{"PageResp", string(data)}, " ")
}

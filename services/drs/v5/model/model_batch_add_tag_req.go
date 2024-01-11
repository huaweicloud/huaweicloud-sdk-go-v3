package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddTagReq 批量添加资源标签
type BatchAddTagReq struct {

	// 标签列表。
	Tags []BatchResourceTag `json:"tags"`
}

func (o BatchAddTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddTagReq struct{}"
	}

	return strings.Join([]string{"BatchAddTagReq", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTagReq 批量删除资源标签
type BatchDeleteTagReq struct {

	// 标签列表。
	Tags []BatchResourceTag `json:"tags"`
}

func (o BatchDeleteTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagReq", string(data)}, " ")
}

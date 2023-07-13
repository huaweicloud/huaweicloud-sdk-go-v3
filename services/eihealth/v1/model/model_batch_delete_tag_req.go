package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTagReq 批量删除镜像tag请求体
type BatchDeleteTagReq struct {

	// 删除镜像tag名称列表
	Tags []string `json:"tags"`
}

func (o BatchDeleteTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagReq", string(data)}, " ")
}

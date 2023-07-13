package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDataReq 删除数据请求体
type BatchDeleteDataReq struct {

	// 删除的数据ID集（项目名称:/路径）
	Paths []string `json:"paths"`
}

func (o BatchDeleteDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDataReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteDataReq", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteLabelReq struct {

	// 批量删除标签id列表
	Ids []string `json:"ids"`
}

func (o BatchDeleteLabelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLabelReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteLabelReq", string(data)}, " ")
}

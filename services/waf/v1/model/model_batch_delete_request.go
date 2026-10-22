package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRequest struct {

	// 要删除的AI应用ID列表
	AppIds []string `json:"app_ids"`
}

func (o BatchDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteRequest", string(data)}, " ")
}

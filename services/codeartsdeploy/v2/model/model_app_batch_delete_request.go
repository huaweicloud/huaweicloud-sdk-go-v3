package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppBatchDeleteRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 应用id列表
	ApplicationIds []string `json:"application_ids"`
}

func (o AppBatchDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppBatchDeleteRequest struct{}"
	}

	return strings.Join([]string{"AppBatchDeleteRequest", string(data)}, " ")
}

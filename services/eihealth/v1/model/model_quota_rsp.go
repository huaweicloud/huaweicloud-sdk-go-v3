package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息
type QuotaRsp struct {

	// 配额项名称，支持USER,PROJECT,USER_PROJECT,STORAGE,PROJECT_APP,PROJECT_NOTEBOOK,PROJECT_WORKFLOW,PROJECT_IMAGE
	Name string `json:"name"`

	// 配额
	Total int32 `json:"total"`

	// 配额项单位
	Unit string `json:"unit"`

	// 配额使用量
	Usage int32 `json:"usage"`
}

func (o QuotaRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaRsp struct{}"
	}

	return strings.Join([]string{"QuotaRsp", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWhitelistRequest struct {

	// 数据库ID列表
	DbIds *[]string `json:"db_ids,omitempty"`

	// 备注
	Desc *string `json:"desc,omitempty"`

	// 是否启用  - true: 启用  - false: 禁用
	Enabled bool `json:"enabled"`
}

func (o UpdateWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhitelistRequest struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistRequest", string(data)}, " ")
}

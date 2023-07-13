package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagParam struct {

	// 环境标签id。
	TagId *int64 `json:"tag_id,omitempty"`

	// 环境标签名称。
	TagName *string `json:"tag_name,omitempty"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 描述信息。
	Descp *string `json:"descp,omitempty"`

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// 环境id列表。
	EnvIdList *[]int64 `json:"env_id_list,omitempty"`

	// 环境标签id列表。
	TagIdList *[]int64 `json:"tag_id_list,omitempty"`

	// 关键字。
	Keyword *string `json:"keyword,omitempty"`

	// 是否分页。
	PageEnable *bool `json:"page_enable,omitempty"`

	// 每页容量。
	PageNumber *int32 `json:"page_number,omitempty"`

	// 当前页码。
	PageSize *int32 `json:"page_size,omitempty"`

	// 新增环境id列表。
	AddEnvIdList *[]int64 `json:"add_env_id_list,omitempty"`

	// 新增环境标签id列表。
	AddTagIdList *[]int64 `json:"add_tag_id_list,omitempty"`

	// 移除环境标签id列表。
	RemoveTagIdList *[]int64 `json:"remove_tag_id_list,omitempty"`

	// 移除的环境id列表。
	RemoveEnvIdList *[]int64 `json:"remove_env_id_list,omitempty"`
}

func (o TagParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagParam struct{}"
	}

	return strings.Join([]string{"TagParam", string(data)}, " ")
}

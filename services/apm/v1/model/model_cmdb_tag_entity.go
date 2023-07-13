package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CmdbTagEntity struct {

	// 环境标签名称。
	Name *string `json:"name,omitempty"`

	// 应用id。
	BusinessId *int64 `json:"business_id,omitempty"`

	// UUID。
	Uuid *string `json:"uuid,omitempty"`

	// 描述信息。
	Descp *string `json:"descp,omitempty"`

	// 创建者id。
	CreatorId *int64 `json:"creator_id,omitempty"`

	// 环境id列表。
	EnvIdList *[]int64 `json:"env_id_list,omitempty"`

	// 环境标签id。
	Id *int64 `json:"id,omitempty"`

	// 创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`

	// 修改时间。
	GmtModify *string `json:"gmt_modify,omitempty"`
}

func (o CmdbTagEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CmdbTagEntity struct{}"
	}

	return strings.Join([]string{"CmdbTagEntity", string(data)}, " ")
}

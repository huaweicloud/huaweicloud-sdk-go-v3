package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CmdbTagEntity struct {

	// 环境标签名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 应用id
	BusinessId *int64 `json:"business_id,omitempty" xml:"business_id"`

	// UUID
	Uuid *string `json:"uuid,omitempty" xml:"uuid"`

	// 描述信息
	Descp *string `json:"descp,omitempty" xml:"descp"`

	// 创建者id
	CreatorId *int64 `json:"creator_id,omitempty" xml:"creator_id"`

	// 环境id列表
	EnvIdList *[]int64 `json:"env_id_list,omitempty" xml:"env_id_list"`

	// 环境标签id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 创建时间
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create"`

	// 修改时间
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify"`
}

func (o CmdbTagEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CmdbTagEntity struct{}"
	}

	return strings.Join([]string{"CmdbTagEntity", string(data)}, " ")
}

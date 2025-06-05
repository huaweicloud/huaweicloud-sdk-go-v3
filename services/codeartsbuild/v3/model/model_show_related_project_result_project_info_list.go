package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowRelatedProjectResultProjectInfoList struct {

	// 唯一标识
	Identifier *string `json:"identifier,omitempty"`

	// 项目名
	Name *string `json:"name,omitempty"`

	// 用户ID
	AuthorId *string `json:"author_id,omitempty"`

	// 是否为项目创建者
	IsCreator *bool `json:"is_creator,omitempty"`

	// 项目创建者租户ID
	AuthorDomainId *string `json:"author_domain_id,omitempty"`
}

func (o ShowRelatedProjectResultProjectInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRelatedProjectResultProjectInfoList struct{}"
	}

	return strings.Join([]string{"ShowRelatedProjectResultProjectInfoList", string(data)}, " ")
}

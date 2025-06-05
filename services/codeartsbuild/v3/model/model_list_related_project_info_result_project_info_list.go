package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListRelatedProjectInfoResultProjectInfoList struct {

	// 唯一标识
	Identifier *string `json:"identifier,omitempty"`

	// 项目名
	Name *string `json:"name,omitempty"`

	// 用户ID
	AuthorId *string `json:"author_id,omitempty"`

	// 状态
	Status *int32 `json:"status,omitempty"`

	// 项目创建者租户ID
	AuthorDomainId *string `json:"author_domain_id,omitempty"`
}

func (o ListRelatedProjectInfoResultProjectInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelatedProjectInfoResultProjectInfoList struct{}"
	}

	return strings.Join([]string{"ListRelatedProjectInfoResultProjectInfoList", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListDomainResponseData struct {

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 域名信息列表
	Records *[]DomainInfo `json:"records,omitempty"`

	// 域名组id
	SetId *string `json:"set_id,omitempty"`

	// 域名总数
	Total *int32 `json:"total,omitempty"`
}

func (o ListDomainResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainResponseData struct{}"
	}

	return strings.Join([]string{"ListDomainResponseData", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListEnterpriseResourceResult struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	ResourceDetail *ListFunctionResult `json:"resource_detail,omitempty"`

	// 标签列表
	Tags *[]KvItem `json:"tags,omitempty"`

	// 系统标签列表
	SysTags *[]KvItem `json:"sys_tags,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ListEnterpriseResourceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseResourceResult struct{}"
	}

	return strings.Join([]string{"ListEnterpriseResourceResult", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationQueryResponseData struct {

	// 应用id
	Id *string `json:"id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用code
	Code *string `json:"code,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 父节点id
	ParentId *string `json:"parent_id,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 应用path路径，由应用id用.拼接
	Path *string `json:"path,omitempty"`
}

func (o ApplicationQueryResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationQueryResponseData struct{}"
	}

	return strings.Join([]string{"ApplicationQueryResponseData", string(data)}, " ")
}

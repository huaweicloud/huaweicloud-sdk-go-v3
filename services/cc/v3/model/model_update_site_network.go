package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteNetwork 更新分支网络的详细信息。
type UpdateSiteNetwork struct {

	// 实例名字。
	Name *string `json:"name,omitempty"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例标签。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdateSiteNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteNetwork struct{}"
	}

	return strings.Join([]string{"UpdateSiteNetwork", string(data)}, " ")
}

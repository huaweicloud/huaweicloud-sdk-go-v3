package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMavenInfoRequest Request Object
type ShowMavenInfoRequest struct {

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// snapshot or releases
	Policy *string `json:"policy,omitempty"`

	// r or rw
	Access *string `json:"access,omitempty"`

	// 是否返回默认仓库 true or false
	Default *string `json:"default,omitempty"`

	// 仓库id 多个仓库id用英文逗号间隔
	Ids *string `json:"ids,omitempty"`
}

func (o ShowMavenInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMavenInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowMavenInfoRequest", string(data)}, " ")
}

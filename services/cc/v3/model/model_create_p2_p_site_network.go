package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateP2PSiteNetwork struct {

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例标签。
	Tags *[]Tag `json:"tags,omitempty"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 端到端(P2P)类型分支网络连接的两个端点定义，长度固定为2的数组。
	Sites []CreateSiteInformation `json:"sites"`
}

func (o CreateP2PSiteNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateP2PSiteNetwork struct{}"
	}

	return strings.Join([]string{"CreateP2PSiteNetwork", string(data)}, " ")
}

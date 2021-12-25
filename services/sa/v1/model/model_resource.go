package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {
	// 资源ID。

	Id string `json:"id"`
	// 资源名称；最大长度255个字符。

	Name string `json:"name"`
	// 资源类型。

	Type string `json:"type"`
	// 云服务名称。

	Provider string `json:"provider"`
	// 区域。

	RegionId *string `json:"region_id,omitempty"`
	// 资源所属租户账号ID。

	DomainId string `json:"domain_id"`
	// 资源所属项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 企业项目ID。

	EpId *string `json:"ep_id,omitempty"`
	// 企业项目名称。

	EpName *string `json:"ep_name,omitempty"`
	// 资源标签 1、最多50个key/values对。 2、values：最大255字符。 3、取值范围：字母数字、空格、“+”、“-”、“=”、“.”、“_”、“:”、“/”、“@”。

	Tags *interface{} `json:"tags,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}

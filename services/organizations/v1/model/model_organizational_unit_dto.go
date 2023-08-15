package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationalUnitDto 包含有关新创建的组织单元的详细信息的结构。
type OrganizationalUnitDto struct {

	// 与组织单元关联的唯一标识符（ID）。
	Id string `json:"id"`

	// 组织单元的统一资源名称。
	Urn string `json:"urn"`

	// 组织单元的名称。
	Name string `json:"name"`

	// 组织单元的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o OrganizationalUnitDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationalUnitDto struct{}"
	}

	return strings.Join([]string{"OrganizationalUnitDto", string(data)}, " ")
}

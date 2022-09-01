package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NeutronSecurityGroup struct {

	// 功能说明：安全组描述 取值范围：0-255个字符
	Description string `json:"description" xml:"description"`

	// 安全组ID
	Id string `json:"id" xml:"id"`

	// 功能说明：安全组名称 取值范围：0-255个字符
	Name string `json:"name" xml:"name"`

	// 安全组规则，详情参见Security Group Rule对象
	SecurityGroupRules []NeutronSecurityGroupRule `json:"security_group_rules" xml:"security_group_rules"`

	// 项目ID
	TenantId string `json:"tenant_id" xml:"tenant_id"`

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 功能说明：资源创建UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	CreatedAt *sdktime.SdkTime `json:"created_at" xml:"created_at"`

	// 功能说明：资源更新UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	UpdatedAt *sdktime.SdkTime `json:"updated_at" xml:"updated_at"`
}

func (o NeutronSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronSecurityGroup struct{}"
	}

	return strings.Join([]string{"NeutronSecurityGroup", string(data)}, " ")
}

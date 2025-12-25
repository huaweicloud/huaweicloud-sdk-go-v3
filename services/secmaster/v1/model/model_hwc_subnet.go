package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcSubnet 云子网
type HwcSubnet struct {

	// 安全组对应的唯一标识
	Id string `json:"id"`

	// 安全组名称
	Name string `json:"name"`

	// 安全组的描述信息
	Description *string `json:"description,omitempty"`

	// 安全组所属的项目ID
	ProjectId string `json:"project_id"`

	// 安全组创建时间 取值范围：UTC时间格式，yyyy-MM-ddTHH:mm:ss
	CreatedAt string `json:"created_at"`

	// 安全组更新时间 取值范围：UTC时间格式，yyyy-MM-ddTHH:mm:ss
	UpdatedAt string `json:"updated_at"`

	// 安全组所属的企业项目ID。 取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 安全组规则
	SecurityGroupRules *[]HwcSubnetSecurityGroupRule `json:"security_group_rules,omitempty"`
}

func (o HwcSubnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcSubnet struct{}"
	}

	return strings.Join([]string{"HwcSubnet", string(data)}, " ")
}

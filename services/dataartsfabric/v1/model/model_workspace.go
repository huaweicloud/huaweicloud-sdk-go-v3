package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Workspace workspace详细信息
type Workspace struct {

	// 工作空间ID。
	Id *string `json:"id,omitempty"`

	// 工作空间名称。
	Name *string `json:"name,omitempty"`

	// 描述。用户输入的描述，最长为255个字符。
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建账号名称。
	CreateDomainName *string `json:"create_domain_name,omitempty"`

	// 创建用户名称。
	CreateUserName *string `json:"create_user_name,omitempty"`

	// Metastore信息，LakeFormation服务的实例Id，即MetaStoreId。
	MetastoreId *string `json:"metastore_id,omitempty"`

	// 访问资源地址
	AccessUrl *string `json:"access_url,omitempty"`

	// 企业项目ID，只有对接了企业项目才可以填写。只能包含字母、数字和中划线，且长度为1到64个字符。默认是0，即default
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o Workspace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Workspace struct{}"
	}

	return strings.Join([]string{"Workspace", string(data)}, " ")
}

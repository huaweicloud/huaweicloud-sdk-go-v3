package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PermissionObject permission列表。
type PermissionObject struct {

	// permission的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// permission列表。 权限格式为： - iam:domain::domain_id。其中： “iam:domain::”为固定格式，“domain_id”为可连接用户的账号ID。 domain_id类型支持输入包括“a~z”、“A~Z”、“0~9”或者“*”，最大长度可以传64。  - organizations:orgPath::org_path。其中： “organizations:orgPath::”为固定格式，org_path为可连接用户的组织路径。 org_path类型支持“a~z”、“A~Z”、“0~9”、“/-*?”或者“*”，最大长度可以传1024。  “*”表示所有终端节点可连接。 示例： - iam:domain::6e9dfd51d1124e8d8498dce894923a0dd - organizations:orgPath::o-3j59d1231uprgk9yuvlidra7zbzfi578/r-rldbu1vmxdw5ahdkknxnvd5rgag77m2z/ou-7tuddd8nh99rebxltawsm6qct5z7rklv/_*
	Permission *string `json:"permission,omitempty"`

	// 终端节点服务白名单类型。  - domainId：基于账户ID配置终端节点服务白名单。  - orgPath：基于账户所在组织路径配置终端节点服务白名单。
	PermissionType *string `json:"permission_type,omitempty"`

	// 白名单的添加时间。 采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o PermissionObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionObject struct{}"
	}

	return strings.Join([]string{"PermissionObject", string(data)}, " ")
}

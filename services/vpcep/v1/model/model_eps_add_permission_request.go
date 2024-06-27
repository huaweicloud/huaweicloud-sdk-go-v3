package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EpsAddPermissionRequest 添加终端节点服务白名单请求体。
type EpsAddPermissionRequest struct {

	// 权限格式为：iam:domain::domain_id或者organizations:orgPath::org_path其中，  - “iam:domain::”和“organizations:orgPath::”为固定格式。  - “domain_id”为可连接用户的账号ID，org_path可连接用户的组织路径 domain_id类型支持输入包括“a~z”、“A~Z”、“0~9”或者“*，最大长度可以传64； org_path类型支持“a~z”、“A~Z”、“0~9”、“/-*?”或者“*”，最大长度可以传1024； 例如：iam:domain::6e9dfd51d1124e8d8498dce894923a0dd或者organizations:orgPath::o-3j59d1231uprgk9yuvlidra7zbzfi578/r-rldbu1vmxdw5ahdkknxnvd5rgag77m2z/ou-7tuddd8nh99rebxltawsm6qct5z7rklv/_*
	Permission string `json:"permission"`

	// 终端节点服务白名单描述
	Description string `json:"description"`
}

func (o EpsAddPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsAddPermissionRequest struct{}"
	}

	return strings.Join([]string{"EpsAddPermissionRequest", string(data)}, " ")
}

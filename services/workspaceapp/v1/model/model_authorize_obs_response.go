package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeObsResponse Response Object
type AuthorizeObsResponse struct {

	// 访问的服务终端节点。
	ServerEndPoint *string `json:"server_end_point,omitempty"`

	// 存放的桶名称。
	BucketName *string `json:"bucket_name,omitempty"`

	// 存放目录。
	Directory *string `json:"directory,omitempty"`

	// 获取的AK。。
	Ak *string `json:"ak,omitempty"`

	// 获取的SK。
	Sk *string `json:"sk,omitempty"`

	// AK/SK和securitytoken的过期时间。。
	ExpiresAt *string `json:"expires_at,omitempty"`

	Policy *Policy `json:"policy,omitempty"`

	// 安全校验token，将所获的AK、SK等信息进行加密后的字符串。
	SecurityToken  *string `json:"security_token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AuthorizeObsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeObsResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeObsResponse", string(data)}, " ")
}

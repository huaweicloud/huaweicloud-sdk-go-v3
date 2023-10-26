package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePermRuleResponse Response Object
type UpdatePermRuleResponse struct {

	// 权限规格的id
	Id *string `json:"id,omitempty"`

	// 授权对象的IP地址或网段
	IpCidr *string `json:"ip_cidr,omitempty"`

	// -| 授权对象的读写权限 rw：默认选项，以读写的方式共享 ro：以只读的方式共享
	RwType *string `json:"rw_type,omitempty"`

	// -| 授权对象的系统用户对文件系统的访问权限。取值如下： no_root_squash：客户端使用的是root用户时，映射到NFS服务器的用户依然为root用户。 root_squash：客户端使用的是root用户时，映射到NFS服务器的用户为NFS的匿名用户（nfsnobody）。 all_squash：默认选项。所有访问NFS服务器的客户端的用户都映射为匿名用户。
	UserType       *string `json:"user_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePermRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdatePermRuleResponse", string(data)}, " ")
}

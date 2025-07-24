package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OnePermRuleRequestInfo 权限规格信息
type OnePermRuleRequestInfo struct {

	// 授权对象的IP地址或网段，设置后不支持修改
	IpCidr string `json:"ip_cidr"`

	// 授权对象的读写权限   - rw：默认选项，以读写的方式共享   - ro：以只读的方式共享   - none: 没有权限
	RwType string `json:"rw_type"`

	// 授权对象的系统用户对文件系统的访问权限。取值如下：  - no_root_squash：默认选项。客户端使用包括root用户在内的任何用户，NFS服务器都保持客户端使用的用户，不做映射。  - root_squash：客户端使用的是root用户时，映射到NFS服务器的用户为NFS的匿名用户（nfsnobody）。客户端使用非root用户时，NFS服务器保持客户端使用的用户，不做映射。  - all_squash：所有访问NFS服务器的客户端的用户都映射为匿名用户。
	UserType string `json:"user_type"`
}

func (o OnePermRuleRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OnePermRuleRequestInfo struct{}"
	}

	return strings.Join([]string{"OnePermRuleRequestInfo", string(data)}, " ")
}

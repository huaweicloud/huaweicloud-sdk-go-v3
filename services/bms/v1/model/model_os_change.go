package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// os-change字段数据结构说明
type OsChange struct {
	// 裸金属服务器管理员帐号的初始登录密码。其中，Linux管理员帐户为root，Windows管理员帐户为Administrator。建议密码复杂度如下：长度为8-26位。密码至少必须包含大写字母、小写字母、数字和特殊字符（!@$%^-_=+[{}]:,./?）中的三种。密码不能包含用户名或用户名的逆序。 说明：对于Windows裸金属服务器，不能包含用户名中超过两个连续字符的部分。对于Linux裸金属服务器也可使用user_data字段实现密码注入，此时adminpass字段无效。adminpass和keyname不能同时有值。adminpass和keyname如果同时为空，此时，metadata中的user_data属性必须有值。

	Adminpass *string `json:"adminpass,omitempty"`
	// 密钥名称。密钥可以通过7.10.3-创建和导入SSH密钥（OpenStack原生）API创建，或者使用7.10.1-查询SSH密钥列表（OpenStack原生）API查询已有的密钥。

	Keyname *string `json:"keyname,omitempty"`
	// 用户ID（登录管理控制台，进入我的凭证，即可看到“用户ID”）。

	Userid *string `json:"userid,omitempty"`
	// 镜像ID。镜像ID可以从镜像服务控制台获取，或者参考《镜像服务API参考》的“查询镜像列表”章节查询。在使用“查询镜像列表”API查询时，可以添加过滤字段“?virtual_env_type=Ironic”来筛选裸金属服务器镜像。

	Imageid string `json:"imageid"`

	Metadata *MetadataInstall `json:"metadata,omitempty"`
}

func (o OsChange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsChange struct{}"
	}

	return strings.Join([]string{"OsChange", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceReq struct {

	// 实例名称。  由英文字符开头，只能由英文字母、数字、中划线组成，长度为4~64的字符。
	Name *string `json:"name,omitempty"`

	// 实例的描述信息。  长度不超过1024的字符串。[且字符串不能包含\">\"与\"<\"，字符串首字符不能为\"=\",\"+\",\"-\",\"@\"的全角和半角字符。](tag:hcs)  > \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\或者\\\"。
	Description *string `json:"description,omitempty"`

	// 安全组ID。  获取方法如下：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// ACL访问控制。
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// 是否开启公网。
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// 实例绑定的弹性IP地址的ID。  以英文逗号隔开多个弹性IP地址的ID。  如果开启了公网访问功能（即enable_publicip为true），该字段为必选。
	PublicipId *string `json:"publicip_id,omitempty"`
}

func (o UpdateInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceReq", string(data)}, " ")
}

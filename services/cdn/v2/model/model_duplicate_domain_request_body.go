package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DuplicateDomainRequestBody **参数解释：** 复制域名请求体 **约束限制：** 不涉及
type DuplicateDomainRequestBody struct {

	// **参数解释：** 存量加速域名，将该域名的配置复制给新添加的域名。  **约束限制：** 已经在CDN添加成功的域名。 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReferenceDomainName string `json:"reference_domain_name"`

	// **参数解释：** 需要添加到CDN控制台加速的域名 > 添加泛域名后，该泛域名所有次级域名均默认接入CDN加速。  **约束限制：** 加速域名不允许重复添加 **取值范围：** - 域名长度不能超过200个字符 - 支持大小写字母、数字、\"-\"、\".\"，首尾字符不能是\"-\"或\".\" - 泛域名场景下支持\"*\"，且\"*\"必须为首字符 - 域名单节点长度不超过63个字符，即：xxx.xxx.com中，xxx的字符数不超过63个字符  **默认取值：** 不涉及
	DomainName string `json:"domain_name"`
}

func (o DuplicateDomainRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DuplicateDomainRequestBody struct{}"
	}

	return strings.Join([]string{"DuplicateDomainRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntranetConnectionModifyRequest **参数解释：** 修改自定义URL请求体。 **约束限制：** 不涉及。  **取值范围：** 不涉及。  **默认取值：** 不涉及。
type IntranetConnectionModifyRequest struct {

	// **参数解释：** 自定义URL，格式为：{协议}://{域名}{路径} **约束限制：** url个数不超过10个，单个url长度不超过1024。 **取值范围：** - 协议范围：http，https，wss，ws。 - 域名范围：支持域名或IP:端口。域名长度不超过63，包含字母、数字、中划线（-)且不能以中划线（-)开头或结尾，顶级域名不能包含数字；端口范围为1-65535。 - 路径范围：斜杠（/）开头，仅包含字母、数字、点号（.）、中划线（-)、下划线（_）、斜杠（/）的路径。 **默认取值：** 不涉及。
	CustomUrls []string `json:"custom_urls"`
}

func (o IntranetConnectionModifyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntranetConnectionModifyRequest struct{}"
	}

	return strings.Join([]string{"IntranetConnectionModifyRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreatePublicZonesTaskRequestBody struct {

	// **参数解释：** 域名。 **约束限制：** 不涉及。 **取值范围：** 由多个以点分隔的字符串组成，可包含字母、数字、汉字、中划线，中划线不能在开头或末尾，单个字符串不超过63个字符，域名总长度不超过254个字符。 **默认取值：** 不涉及。
	ZoneNames []string `json:"zone_names"`

	// **参数解释：** 域名所属的企业项目ID。可以使用该字段过滤企业项目下的域名。 **约束限制：** 不涉及。 **取值范围：** 最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。 **默认取值：** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o BatchCreatePublicZonesTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicZonesTaskRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicZonesTaskRequestBody", string(data)}, " ")
}

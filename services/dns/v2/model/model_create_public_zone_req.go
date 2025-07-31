package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublicZoneReq **参数解释：** 创建公网域名请求。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type CreatePublicZoneReq struct {

	// **参数解释：** 域名。 **约束限制：** 不涉及。 **取值范围：** 由多个以点分隔的字符串组成，可包含字母、数字、汉字、中划线，中划线不能在开头或末尾，单个字符串不超过63个字符，域名总长度不超过254个字符。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 域名的描述信息。 **约束限制：** 不涉及。 **取值范围：** 长度不超过255个字符。 **默认取值：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 域名类型。 **约束限制：** 不涉及。 **取值范围：** public：公网域名 **默认取值：** 不涉及。
	ZoneType *string `json:"zone_type,omitempty"`

	// **参数解释：** 管理该域名的管理员邮箱，用于生成该域名的SOA记录。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Email *string `json:"email,omitempty"`

	// **参数解释：** 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。 **约束限制：** 不涉及。 **取值范围：** 1~2147483647。 **默认取值：** 300
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 域名所属的企业项目ID。可以使用该字段过滤企业项目下的域名。 **约束限制：** 不涉及。 **取值范围：** 最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 资源标签。 **取值范围：** 不涉及。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreatePublicZoneReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneReq struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneReq", string(data)}, " ")
}

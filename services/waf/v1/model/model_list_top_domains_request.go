package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopDomainsRequest Request Object
type ListTopDomainsRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符  **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 起始时间（13位毫秒时间戳），需要和to同时使用。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	From int64 `json:"from"`

	// **参数解释：** 结束时间（13位毫秒时间戳），需要和from同时使用。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	To int64 `json:"to"`

	// **参数解释：** 查询受攻击次数排名在前几的结果 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 5
	Top *int32 `json:"top,omitempty"`

	// **参数解释：** 域名ID，查询的域名列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *[]string `json:"hosts,omitempty"`
}

func (o ListTopDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListTopDomainsRequest", string(data)}, " ")
}

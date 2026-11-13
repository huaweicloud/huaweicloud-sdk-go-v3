package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationDetailsResponse Response Object
type ListApplicationDetailsResponse struct {
	BasicInfo *ApplicationBasicInfoRsp `json:"basic_info,omitempty"`

	SystemConfig *SysconfigInfo `json:"system_config,omitempty"`

	// **参数解释**： 文档块信息，仅当query_type为2或者4时返回。 支持more和basic两个列表，分别标识基础文档块和扩展文档块。 - basic：基础文档块列表，元素为文档块对象。 - more：扩展文档块列表，元素为文档块对象或分组标题对象。 - release_state：发布状态，true为已发布，false为未发布。 文档块对象支持字段如下： - code：文档块编码。 - pinyin：文档块简称。 - quanpin：文档块全程。 - name：文档块名称。 - icon：文档块图标。 - show：是否展示： true为展示，false为不展示。 - source：文档块来源，0标识系统文档块，分组标题对象中1标识自定义分组。 - darkIcon：暗黑模式文档块图标。 - key：文档块key值。 - data：文档块附加数据，可为字符串或对象（可选）。 分组标题对象字段如下： - source：分组来源。 - title：分组标题。 **取值范围**： 不涉及。
	BlocksInfo *interface{} `json:"blocks_info,omitempty"`

	// **参数解释**： 发布信息，仅当quey_type为3或者4时返回。 **取值范围**： 不涉及。
	ReleaseInfo *[]ReleaseAppInfo `json:"release_info,omitempty"`

	IdpInfo *IdpDetailInfo `json:"idp_info,omitempty"`

	// **参数解释**： 发布信息的总数量，仅当查询包含发布信息时，该字段返回。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	XTraceId       *string `json:"X-Trace-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListApplicationDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationDetailsResponse", string(data)}, " ")
}

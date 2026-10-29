package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCriterionSetsByLanguageRequest Request Object
type ListCriterionSetsByLanguageRequest struct {

	// **参数解释:** 规则集语言。 **约束限制:** 不涉及。 **取值范围:** 当前支持的语言有： - ARKTS - C# - CANGJIE - C++ - CSS - GO - HTML - JAVA - JAVASCRIPT - KOTLIN - LUA - PHP - PYTHON - RUST - SCALA - SHELL - SQL - TERRAFORM - TYPESCRIPT - AZURERESOURCEMANAGER - DOCKER - KUBERNETES **默认取值：** 不涉及。
	Language string `json:"language"`

	// **参数解释:** 搜索关键字。 **约束限制:** 不涉及。 **取值范围:** 不涉及。 **默认取值：** 不涉及。
	Search *string `json:"search,omitempty"`

	// **参数解释:** 每页规则集数量。 **约束限制:** 不涉及。 **取值范围:** 1~1000 **默认取值:** 1000
	PageSize *int32 `json:"page_size,omitempty"`

	// **参数解释:** 规则集页数。 **约束限制:** 不涉及。 **取值范围:** 0-999 **默认取值:** 1
	Page *int32 `json:"page,omitempty"`

	// **参数解释:** 项目的32位UUID，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/zh-cn/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)接口获取。 **约束限制:** 不涉及。 **取值范围:** 32位UUID。 **默认取值:** 不涉及。
	ProjectId string `json:"project_id"`

	// 是否可信
	IsReliable *string `json:"is_reliable,omitempty"`
}

func (o ListCriterionSetsByLanguageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCriterionSetsByLanguageRequest struct{}"
	}

	return strings.Join([]string{"ListCriterionSetsByLanguageRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTasksListRequest Request Object
type ShowTasksListRequest struct {

	// **参数解释：** 分页页码。 **约束限制：** 分页页码≥1。 **取值范围:** 1-2147483647 **默认取值：**  不涉及。
	Page *int32 `json:"page,omitempty"`

	// **参数解释：** 分页大小。 **约束限制：** 数字1-200之间。 **取值范围:** 1-200 **默认取值：** 不涉及。
	PageSize *int32 `json:"page_size,omitempty"`

	// **参数解释：** 需要搜索的任务名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 项目的32位UUID，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/zh-cn/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)接口获取。 **约束条件：** 不涉及。 **取值范围：** 长度为[1~32]个字符。 **默认取值：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 门禁结果。 **约束限制：** 不涉及。 **取值范围：** - success：门禁通过。 - error：门禁不通过。 **默认取值：** 不涉及。
	ReviewResultTypes *string `json:"review_result_types,omitempty"`

	// **参数解释：** 检查状态。 **约束限制：** 不涉及。 **取值范围：** - success：检查成功。 - failed：检查失败。 - aborted：检查终止。 - executing：检查中。 - never_execute：未检查。 **默认取值：** 不涉及。
	JobStatusTypes *string `json:"job_status_types,omitempty"`

	// **参数解释：** 是否按照执行时间排序。 **约束限制：** 不涉及。 **取值范围：** - up：按照执行时间升序排序。 - down：按照执行时间降序排序。 **默认取值：** 不涉及。
	OrderByExecuteTime *string `json:"order_by_execute_time,omitempty"`

	// **参数解释：** 任务的分组ID，可通过列出项目下的分组信息接口获取任务的分组ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TaskClassificationId *string `json:"task_classification_id,omitempty"`

	// 列表类接口参数
	Offset *int32 `json:"offset,omitempty"`

	// 列表类接口参数
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序字段,任务列表每次只能根据一个字段进行排序。 **约束限制：** 不涉及。
	SortField *string `json:"sort_field,omitempty"`

	// **参数解释：** 排序顺序，取值范围：\"up\"（升序）、\"down\"（降序） **约束限制：** 不涉及。
	SortOrder *string `json:"sort_order,omitempty"`

	// **参数解释：** 操作人，填写账号用户名。 **约束限制：** 不涉及。 **取值范围：** 长度为[1~128]个字符。 **默认取值：** 不涉及。
	Operator *string `json:"operator,omitempty"`

	// **参数解释：** 代码仓地址。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RepoUrl *string `json:"repo_url,omitempty"`

	// **参数解释：** 分支。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Branch *string `json:"branch,omitempty"`

	// **参数解释：** 和repo_url配合使用，查询任务时是否完全匹配repo_url。 **约束限制：** 不涉及。 **取值范围：** - true：只查询代码仓地址为repo_url的任务。 - false：查询ssh协议的代码仓地址或http协议的代码仓地址和repo_url匹配的任务。 **默认取值：** false。
	RepoUrlFullMatch *string `json:"repo_url_full_match,omitempty"`
}

func (o ShowTasksListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasksListRequest struct{}"
	}

	return strings.Join([]string{"ShowTasksListRequest", string(data)}, " ")
}

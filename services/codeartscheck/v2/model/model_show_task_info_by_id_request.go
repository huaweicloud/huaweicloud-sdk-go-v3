package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskInfoByIdRequest Request Object
type ShowTaskInfoByIdRequest struct {

	// **参数解释：** 任务ID，即新建检查任务接口的返回值，每个代码检查任务及其下的分支任务、通过API创建的增量任务都有唯一的任务ID。即[[查询任务列表](https://support.huaweicloud.com/api-codecheck/ShowTasksList.html)](tag:hws)[[查询任务列表](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ShowTasksList.html)](tag:hws_hk)[查询任务列表](tag:hcs,hcs_sm)接口响应中的id字段。 **约束限制：** 不涉及。 **取值范围：** 长度为[1~32]个字符。 **默认取值：** 不涉及。
	TaskId string `json:"task_id"`

	// **参数解释：** 合并请求ID。 **约束限制：** 不涉及。 **取值范围：** 长度为[1~128]个字符。 **默认取值：** 不涉及。
	MergeId *string `json:"merge_id,omitempty"`

	// **参数解释：** 执行记录ID，即每一次执行版本级或门禁级扫描的唯一ID。通过[[执行检查任务](https://support.huaweicloud.com/api-codecheck/RunTaskV2.html)](tag:hws)[[执行检查任务](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/RunTaskV2.html)](tag:hws_hk)[执行检查任务](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)接口的返回值可获取执行记录ID。 **约束限制：** 不涉及。 **取值范围：** 长度为[1~64]个字符。 **默认取值：** 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释：** 操作者。 **约束限制：** 不涉及。 **取值范围：** 长度为[1~128]个字符。 **默认取值：** 不涉及。
	Operator *string `json:"operator,omitempty"`
}

func (o ShowTaskInfoByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInfoByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskInfoByIdRequest", string(data)}, " ")
}

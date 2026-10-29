package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeGateRequestBody **参数解释：** 门禁配置。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type UpdateCodeGateRequestBody struct {

	// **参数解释：** 任务门禁配置列表。 **约束限制：** 不涉及。
	ReviewData []ReviewDataInfoRequestParam `json:"reviewData"`

	// **参数解释：** 任务ID，即新建检查任务接口的返回值，每个代码检查任务及其下的分支任务、通过API创建的增量任务都有唯一的任务ID。即[[查询任务列表](https://support.huaweicloud.com/api-codecheck/ShowTasksList.html)](tag:hws)[[查询任务列表](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ShowTasksList.html)](tag:hws_hk)[查询任务列表](tag:hcs,hcs_sm)接口响应中的id字段。 **约束限制：** 不涉及。 **取值范围：** 长度为[1~32]个字符。 **默认取值：** 不涉及。
	Id string `json:"id"`
}

func (o UpdateCodeGateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeGateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCodeGateRequestBody", string(data)}, " ")
}

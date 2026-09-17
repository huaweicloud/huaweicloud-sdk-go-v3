package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlanRequest Request Object
type ListPlanRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// **参数解释：** 发布/迭代名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyWord *string `json:"key_word,omitempty"`

	// **参数解释：** 更新发布/迭代时间，unix时间戳，单位：毫秒  样例：1576114296000,1576114396000 **约束限制：**  起止时间均为13位的时间戳字符串，使用英文逗号分割。 **取值范围：** 不涉及 **默认取值：** 不涉及
	UpdatedTimeInterval *string `json:"updated_time_interval,omitempty"`
}

func (o ListPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlanRequest struct{}"
	}

	return strings.Join([]string{"ListPlanRequest", string(data)}, " ")
}

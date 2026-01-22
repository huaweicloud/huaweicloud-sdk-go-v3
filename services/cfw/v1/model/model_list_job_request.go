package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobRequest Request Object
type ListJobRequest struct {

	// **参数解释**： 创建按需防火墙返回的任务ID，可通过调用[创建防火墙接口](CreateFirewall.xml)返回值获得。返回值中job_id即为此处的job_id **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	JobId string `json:"job_id"`
}

func (o ListJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobRequest struct{}"
	}

	return strings.Join([]string{"ListJobRequest", string(data)}, " ")
}

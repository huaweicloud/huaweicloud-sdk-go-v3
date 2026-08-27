package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EvolveTaskBatchDeleteReq struct {

	// **参数解释**： 演化任务完成时间,单位毫秒。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、下划线和中划线，长度为[1-128]个字符。最多支持设置50个作业ID。 **默认取值**： 不涉及
	Resources []string `json:"resources"`
}

func (o EvolveTaskBatchDeleteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvolveTaskBatchDeleteReq struct{}"
	}

	return strings.Join([]string{"EvolveTaskBatchDeleteReq", string(data)}, " ")
}

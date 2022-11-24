package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数工作流URN, 格式为： urn:fss:<region_id>:<project_id>:workflow:\\<package\\>:<workflow_name>:\\<version\\> 注意： package当前只支持default version当前只支持latest
type WorkflowUrn struct {
}

func (o WorkflowUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowUrn struct{}"
	}

	return strings.Join([]string{"WorkflowUrn", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryTemplate 模板中构建执行步骤列表
type QueryTemplate struct {

	// 构建执行的步骤
	Steps []CreateBuildJobSteps `json:"steps"`
}

func (o QueryTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTemplate struct{}"
	}

	return strings.Join([]string{"QueryTemplate", string(data)}, " ")
}

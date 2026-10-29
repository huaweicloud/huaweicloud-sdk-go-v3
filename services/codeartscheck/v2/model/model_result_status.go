package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultStatus **参数解释：** 接口响应状态。 **取值范围：** - success：接口响应成功。 - error：接口响应失败。
type ResultStatus struct {
}

func (o ResultStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultStatus struct{}"
	}

	return strings.Join([]string{"ResultStatus", string(data)}, " ")
}

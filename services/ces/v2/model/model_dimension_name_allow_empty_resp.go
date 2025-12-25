package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimensionNameAllowEmptyResp **参数解释**： 资源维度名称，各服务资源的维度名称可查看：“[服务维度名称](ces_03_0059.xml)”。 **取值范围**： 必须以字母开头，多维度用\",\"分隔，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32。目前最大支持4个维度。举例：单维度场景：instance_id；多维度场景：instance_id,disk
type DimensionNameAllowEmptyResp struct {
}

func (o DimensionNameAllowEmptyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionNameAllowEmptyResp struct{}"
	}

	return strings.Join([]string{"DimensionNameAllowEmptyResp", string(data)}, " ")
}

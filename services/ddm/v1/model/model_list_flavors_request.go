package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {
	// 引擎id,通过查询DDM引擎信息接口获取。

	EngineId string `json:"engine_id"`
	// 分页参数：起始值 [大于等于0] 。默认值是0。

	Offset *int32 `json:"offset,omitempty"`
	// 分页参数：每页多少条 [大于0且小于等于128]。默认值是128。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateDbDataReq **参数解释**： 校验schema或table列表数据请求体。 **取值范围**： 不涉及。
type ValidateDbDataReq struct {

	// **参数解释**： 类型。 **取值范围**： - schema - table
	Type string `json:"type"`

	// schema或者table列表。
	Data []string `json:"data"`
}

func (o ValidateDbDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDbDataReq struct{}"
	}

	return strings.Join([]string{"ValidateDbDataReq", string(data)}, " ")
}

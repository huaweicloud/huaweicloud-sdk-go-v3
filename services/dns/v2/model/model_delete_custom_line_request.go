package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteCustomLineRequest struct {
	// 解析线路ID。

	LineId string `json:"line_id"`
}

func (o DeleteCustomLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomLineRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomLineRequest", string(data)}, " ")
}

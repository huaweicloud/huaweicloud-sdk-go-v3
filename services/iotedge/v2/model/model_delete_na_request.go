package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteNaRequest struct {
	// 北向数据接收端点ID

	NaId string `json:"na_id"`
}

func (o DeleteNaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNaRequest struct{}"
	}

	return strings.Join([]string{"DeleteNaRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowNaRequest struct {

	// 北向数据接收端点ID
	NaId string `json:"na_id"`
}

func (o ShowNaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNaRequest struct{}"
	}

	return strings.Join([]string{"ShowNaRequest", string(data)}, " ")
}

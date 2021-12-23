package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateNaRequest struct {
	// 北向数据接收端点ID

	NaId string `json:"na_id"`

	Body *UpdateNaRequestDto `json:"body,omitempty"`
}

func (o UpdateNaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNaRequest struct{}"
	}

	return strings.Join([]string{"UpdateNaRequest", string(data)}, " ")
}

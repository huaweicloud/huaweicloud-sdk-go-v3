package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTuningRequest Request Object
type CreateTuningRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateTuningReq `json:"body,omitempty"`
}

func (o CreateTuningRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTuningRequest struct{}"
	}

	return strings.Join([]string{"CreateTuningRequest", string(data)}, " ")
}

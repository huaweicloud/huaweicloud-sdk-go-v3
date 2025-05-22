package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGeminiDbDualActiveResponse Response Object
type DeleteGeminiDbDualActiveResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGeminiDbDualActiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGeminiDbDualActiveResponse struct{}"
	}

	return strings.Join([]string{"DeleteGeminiDbDualActiveResponse", string(data)}, " ")
}

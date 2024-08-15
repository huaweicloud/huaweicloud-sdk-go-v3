package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetScriptJobInfoResponse Response Object
type GetScriptJobInfoResponse struct {
	Data           *JobScriptOrderInfoModel `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o GetScriptJobInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetScriptJobInfoResponse struct{}"
	}

	return strings.Join([]string{"GetScriptJobInfoResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCheckResultResponse Response Object
type RunCheckResultResponse struct {
	Result         *CheckResultResultBody `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o RunCheckResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCheckResultResponse struct{}"
	}

	return strings.Join([]string{"RunCheckResultResponse", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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

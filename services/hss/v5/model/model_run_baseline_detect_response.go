package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunBaselineDetectResponse Response Object
type RunBaselineDetectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RunBaselineDetectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunBaselineDetectResponse struct{}"
	}

	return strings.Join([]string{"RunBaselineDetectResponse", string(data)}, " ")
}

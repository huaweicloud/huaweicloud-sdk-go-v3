package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPatchBaselineRequest Request Object
type ShowPatchBaselineRequest struct {

	// 基线id
	BaselineId string `json:"baseline_id"`
}

func (o ShowPatchBaselineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPatchBaselineRequest struct{}"
	}

	return strings.Join([]string{"ShowPatchBaselineRequest", string(data)}, " ")
}

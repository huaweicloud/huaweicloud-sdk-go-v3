package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTargetPasswordRequest Request Object
type ShowTargetPasswordRequest struct {

	// 模板的ID
	Id string `json:"id"`
}

func (o ShowTargetPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTargetPasswordRequest struct{}"
	}

	return strings.Join([]string{"ShowTargetPasswordRequest", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyApplicationRequest Request Object
type ModifyApplicationRequest struct {

	// 应用ID。
	AppId string `json:"app_id"`

	Body *ModifyApplicationRequestBody `json:"body,omitempty"`
}

func (o ModifyApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyApplicationRequest struct{}"
	}

	return strings.Join([]string{"ModifyApplicationRequest", string(data)}, " ")
}

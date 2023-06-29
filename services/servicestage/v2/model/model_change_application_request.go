package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeApplicationRequest Request Object
type ChangeApplicationRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id"`

	Body *ApplicationModify `json:"body,omitempty"`
}

func (o ChangeApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeApplicationRequest struct{}"
	}

	return strings.Join([]string{"ChangeApplicationRequest", string(data)}, " ")
}

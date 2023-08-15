package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateRequest Request Object
type UpdateTemplateRequest struct {

	// 需要修改信息的模板的ID
	Id string `json:"id"`

	Body *UpdateTemplateReq `json:"body,omitempty"`
}

func (o UpdateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateRequest", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTempRequest Request Object
type ShowTempRequest struct {

	// 事务id
	TemplateId int32 `json:"template_id"`
}

func (o ShowTempRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTempRequest struct{}"
	}

	return strings.Join([]string{"ShowTempRequest", string(data)}, " ")
}

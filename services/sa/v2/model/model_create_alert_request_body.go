package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRequestBody CreateAlertRequestBody
type CreateAlertRequestBody struct {
	DataObject *CreateAlert `json:"data_object,omitempty"`
}

func (o CreateAlertRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlertRequestBody", string(data)}, " ")
}

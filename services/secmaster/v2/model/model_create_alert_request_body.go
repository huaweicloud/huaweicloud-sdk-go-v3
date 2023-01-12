package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRequestBody
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

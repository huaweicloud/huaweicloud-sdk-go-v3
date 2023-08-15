package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAlertRequestBody ChangeAlertRequestBody
type ChangeAlertRequestBody struct {
	DataObject *Alert `json:"data_object,omitempty"`
}

func (o ChangeAlertRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAlertRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeAlertRequestBody", string(data)}, " ")
}

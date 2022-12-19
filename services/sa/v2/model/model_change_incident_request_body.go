package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIncidentRequestBody
type ChangeIncidentRequestBody struct {
	DataObject *Incident `json:"data_object,omitempty"`
}

func (o ChangeIncidentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIncidentRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeIncidentRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeResourceRequestBody ChangeResourceRequestBody
type ChangeResourceRequestBody struct {
	DataObject *ResourceDetail `json:"data_object"`
}

func (o ChangeResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeResourceRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeResourceRequestBody", string(data)}, " ")
}

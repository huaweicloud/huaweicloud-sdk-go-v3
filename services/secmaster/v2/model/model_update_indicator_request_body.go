package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIndicatorRequestBody update indicator request body
type UpdateIndicatorRequestBody struct {
	DataObject *IndicatorDataObjectDetail `json:"data_object,omitempty"`
}

func (o UpdateIndicatorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIndicatorRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIndicatorRequestBody", string(data)}, " ")
}

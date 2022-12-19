package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// indicator batch operation response
type IndicatorBatchOperateResponse struct {

	// id list
	SuccessIds []string `json:"success_ids"`

	// id list
	ErrorIds *[]string `json:"error_ids,omitempty"`
}

func (o IndicatorBatchOperateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorBatchOperateResponse struct{}"
	}

	return strings.Join([]string{"IndicatorBatchOperateResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// delete indicator request body
type DeleteIndicatorRequestBody struct {

	// id list
	BatchIds *[]string `json:"batch_ids,omitempty"`
}

func (o DeleteIndicatorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIndicatorRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteIndicatorRequestBody", string(data)}, " ")
}

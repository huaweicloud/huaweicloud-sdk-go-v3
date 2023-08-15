package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableRequest Request Object
type DeleteTableRequest struct {

	// 表ID
	TableId string `json:"table_id"`
}

func (o DeleteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableRequest struct{}"
	}

	return strings.Join([]string{"DeleteTableRequest", string(data)}, " ")
}

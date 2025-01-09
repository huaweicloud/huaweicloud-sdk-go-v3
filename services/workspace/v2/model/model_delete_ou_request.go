package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOuRequest Request Object
type DeleteOuRequest struct {

	// OU的id。
	OuId string `json:"ou_id"`
}

func (o DeleteOuRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOuRequest struct{}"
	}

	return strings.Join([]string{"DeleteOuRequest", string(data)}, " ")
}

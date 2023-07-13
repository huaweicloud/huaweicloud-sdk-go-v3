package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAkSkRequest Request Object
type DeleteAkSkRequest struct {

	// 需要删除的ak信息。
	Ak string `json:"ak"`
}

func (o DeleteAkSkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAkSkRequest struct{}"
	}

	return strings.Join([]string{"DeleteAkSkRequest", string(data)}, " ")
}

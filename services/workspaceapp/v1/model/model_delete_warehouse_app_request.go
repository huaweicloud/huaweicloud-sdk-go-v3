package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWarehouseAppRequest Request Object
type DeleteWarehouseAppRequest struct {

	// 应用仓库中的应用记录ID。
	Id string `json:"id"`
}

func (o DeleteWarehouseAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWarehouseAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteWarehouseAppRequest", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDataStoreRequest struct {
	// 存储 ID

	DataStoreId string `json:"data_store_id"`
}

func (o DeleteDataStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataStoreRequest struct{}"
	}

	return strings.Join([]string{"DeleteDataStoreRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DownloadProductsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 待导出产品ID列表
	ProductIds *[]int32 `json:"product_ids,omitempty"`
}

func (o DownloadProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadProductsRequest struct{}"
	}

	return strings.Join([]string{"DownloadProductsRequest", string(data)}, " ")
}

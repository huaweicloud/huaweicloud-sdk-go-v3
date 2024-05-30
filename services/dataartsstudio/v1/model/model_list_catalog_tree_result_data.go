package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogTreeResultData data，统一的返回结果的最外层数据结构。
type ListCatalogTreeResultData struct {

	// value，统一的返回结果的外层数据结构。
	Value *[]BizCatalogVo `json:"value,omitempty"`
}

func (o ListCatalogTreeResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogTreeResultData struct{}"
	}

	return strings.Join([]string{"ListCatalogTreeResultData", string(data)}, " ")
}

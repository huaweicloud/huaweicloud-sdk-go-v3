package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCatalogResultData data，统一的返回结果的最外层数据结构。
type CreateCatalogResultData struct {
	Value *BizCatalogVo `json:"value,omitempty"`
}

func (o CreateCatalogResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCatalogResultData struct{}"
	}

	return strings.Join([]string{"CreateCatalogResultData", string(data)}, " ")
}

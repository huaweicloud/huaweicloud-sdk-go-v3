package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProtocolMappingsRequest struct {
	// 设备关联的产品ID，用于唯一标识一个产品模型，在管理门户导入产品模型后由平台分配获得。

	ProductId string `json:"product_id"`
}

func (o ShowProtocolMappingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtocolMappingsRequest struct{}"
	}

	return strings.Join([]string{"ShowProtocolMappingsRequest", string(data)}, " ")
}

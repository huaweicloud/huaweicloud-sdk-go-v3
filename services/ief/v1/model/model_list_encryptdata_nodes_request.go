package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEncryptdataNodesRequest struct {

	// 加密数据ID
	EncryptdataId string `json:"encryptdata_id"`

	// 查询返回记录的数量限制
	Limit *int32 `json:"limit,omitempty"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListEncryptdataNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEncryptdataNodesRequest struct{}"
	}

	return strings.Join([]string{"ListEncryptdataNodesRequest", string(data)}, " ")
}

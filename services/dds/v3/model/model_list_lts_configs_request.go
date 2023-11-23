package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsConfigsRequest Request Object
type ListLtsConfigsRequest struct {

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。取值范围：0~50。不传该参数时，默认查询前50条实例的云服务日志配置信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListLtsConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListLtsConfigsRequest", string(data)}, " ")
}

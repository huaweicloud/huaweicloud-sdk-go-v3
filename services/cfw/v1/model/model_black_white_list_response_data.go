package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BlackWhiteListResponseData 查询黑白名单返回值
type BlackWhiteListResponseData struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
	ObjectId *string `json:"object_id,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 查询的总条数
	Total *int32 `json:"total,omitempty"`

	// 黑白名单记录
	Records *[]BlackWhiteListResponseDataRecords `json:"records,omitempty"`
}

func (o BlackWhiteListResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlackWhiteListResponseData struct{}"
	}

	return strings.Join([]string{"BlackWhiteListResponseData", string(data)}, " ")
}

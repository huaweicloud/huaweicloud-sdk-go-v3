package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowLinkResponse struct {

	// 连接列表，请参见links数据结构说明
	Links *[]Links `json:"links,omitempty" xml:"links"`

	// 表/文件迁移不支持哪些数据源迁移到哪些数据源
	FromToUnMapping *string `json:"fromTo-unMapping,omitempty" xml:"fromTo-unMapping"`

	// 整库迁移支持哪些数据源迁移到哪些数据源
	BatchFromToMapping *string `json:"batchFromTo-mapping,omitempty" xml:"batchFromTo-mapping"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowLinkResponse", string(data)}, " ")
}

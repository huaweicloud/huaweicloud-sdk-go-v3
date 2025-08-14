package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJarPackageHostInfoResponse Response Object
type ListJarPackageHostInfoResponse struct {

	// **参数解释** 包含该中间件的服务器总数 **取值范围**: 字符长度0-10000位
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 服务器列表 **取值范围**: 不涉及
	DataList       *[]JarPackageHostInfo `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListJarPackageHostInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJarPackageHostInfoResponse struct{}"
	}

	return strings.Join([]string{"ListJarPackageHostInfoResponse", string(data)}, " ")
}

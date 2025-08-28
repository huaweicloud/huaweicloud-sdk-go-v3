package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJarPackageInfoResponse Response Object
type ListJarPackageInfoResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 中间件列表 **取值范围**: 最小值0，最大值10000
	DataList       *[]JarPackageInfo `json:"data_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListJarPackageInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJarPackageInfoResponse struct{}"
	}

	return strings.Join([]string{"ListJarPackageInfoResponse", string(data)}, " ")
}

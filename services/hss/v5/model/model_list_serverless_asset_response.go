package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerlessAssetResponse Response Object
type ListServerlessAssetResponse struct {

	// **参数解释**: 容器总数 **取值范围**: 取值0-2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 容器基本信息列表 **取值范围**: 取值0-2147483647个资产对象
	DataList       *[]ServerlessAssetBaseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListServerlessAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerlessAssetResponse struct{}"
	}

	return strings.Join([]string{"ListServerlessAssetResponse", string(data)}, " ")
}

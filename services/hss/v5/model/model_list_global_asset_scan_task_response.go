package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalAssetScanTaskResponse Response Object
type ListGlobalAssetScanTaskResponse struct {

	// **参数解释** 是否存在全量扫描任务 **取值范围** - true：是 - false：否
	Exist          *bool `json:"exist,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListGlobalAssetScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalAssetScanTaskResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalAssetScanTaskResponse", string(data)}, " ")
}

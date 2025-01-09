package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAzResponse Response Object
type ListAzResponse struct {

	// 云应用支持的可用分区表格，按站点分类。
	Azs map[string][]AvailabilityZoneInfo `json:"azs,omitempty"`

	// 默认站点类型。
	DefaultType *string `json:"default_type,omitempty"`

	// 云应用支持的站点类型。
	SupportType    *[]string `json:"support_type,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAzResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAzResponse struct{}"
	}

	return strings.Join([]string{"ListAzResponse", string(data)}, " ")
}

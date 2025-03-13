package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityNoMaskingTableResultResponse Response Object
type ShowSecurityNoMaskingTableResultResponse struct {

	// 表总数
	Total *int32 `json:"total,omitempty"`

	// 查询的表集合
	Tables         *[]DiagnoseNoMaskingDetail `json:"tables,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowSecurityNoMaskingTableResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityNoMaskingTableResultResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityNoMaskingTableResultResponse", string(data)}, " ")
}

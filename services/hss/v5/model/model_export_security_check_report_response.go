package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSecurityCheckReportResponse Response Object
type ExportSecurityCheckReportResponse struct {

	// 配置检测信息总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 数据列表
	DataList       *[]ExportSecurityCheckInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ExportSecurityCheckReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSecurityCheckReportResponse struct{}"
	}

	return strings.Join([]string{"ExportSecurityCheckReportResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWdrReportResponse Response Object
type CreateWdrReportResponse struct {

	// 操作结果
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateWdrReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWdrReportResponse struct{}"
	}

	return strings.Join([]string{"CreateWdrReportResponse", string(data)}, " ")
}

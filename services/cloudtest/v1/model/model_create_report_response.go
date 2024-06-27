package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportResponse Response Object
type CreateReportResponse struct {
	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportResponse struct{}"
	}

	return strings.Join([]string{"CreateReportResponse", string(data)}, " ")
}

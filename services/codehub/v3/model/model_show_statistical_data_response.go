package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStatisticalDataResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RepositoryStatisticsVo `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStatisticalDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticalDataResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticalDataResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRepositoryStatisticsResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *RepoStatisticsLaunch `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepositoryStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryStatisticsResponse", string(data)}, " ")
}

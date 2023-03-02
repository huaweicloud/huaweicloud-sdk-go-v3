package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSpecIssueStayTimesResponse struct {

	// 计算失败的工作项id,一般指未关闭的工作项
	Fails *[]string `json:"fails,omitempty"`

	// 工作项时间数据（创建时间-关闭时间）
	Data *[]ListSpecIssueStayTimesResponseBodyData `json:"data,omitempty"`

	// 停留时间求和（单位：秒）
	TotalStayTime *int64 `json:"total_stay_time,omitempty"`

	// 停留时间求和的工作项个数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSpecIssueStayTimesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecIssueStayTimesResponse struct{}"
	}

	return strings.Join([]string{"ListSpecIssueStayTimesResponse", string(data)}, " ")
}

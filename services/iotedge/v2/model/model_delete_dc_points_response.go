package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDcPointsResponse Response Object
type DeleteDcPointsResponse struct {
	Success *[]string `json:"success,omitempty"`

	// 删除失败的点位和详情
	Failure        *[]DeleteDcPointsFailedDetail `json:"failure,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o DeleteDcPointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDcPointsResponse struct{}"
	}

	return strings.Join([]string{"DeleteDcPointsResponse", string(data)}, " ")
}

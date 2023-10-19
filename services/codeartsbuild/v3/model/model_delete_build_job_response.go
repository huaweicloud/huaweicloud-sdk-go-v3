package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBuildJobResponse Response Object
type DeleteBuildJobResponse struct {
	Result *DeleteBuildJobResponseBodyResult `json:"result,omitempty"`

	// 状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBuildJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBuildJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteBuildJobResponse", string(data)}, " ")
}

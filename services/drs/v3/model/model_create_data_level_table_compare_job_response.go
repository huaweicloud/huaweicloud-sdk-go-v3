package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataLevelTableCompareJobResponse Response Object
type CreateDataLevelTableCompareJobResponse struct {

	// 对比任务ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataLevelTableCompareJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataLevelTableCompareJobResponse struct{}"
	}

	return strings.Join([]string{"CreateDataLevelTableCompareJobResponse", string(data)}, " ")
}

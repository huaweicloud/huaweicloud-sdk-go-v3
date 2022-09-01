package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSqlFilterControlResponse struct {

	// 开启/关闭SQL限流任务ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSqlFilterControlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlFilterControlResponse struct{}"
	}

	return strings.Join([]string{"UpdateSqlFilterControlResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRayJobResponse Response Object
type CancelRayJobResponse struct {

	// **参数解释**：作业ID。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelRayJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRayJobResponse struct{}"
	}

	return strings.Join([]string{"CancelRayJobResponse", string(data)}, " ")
}

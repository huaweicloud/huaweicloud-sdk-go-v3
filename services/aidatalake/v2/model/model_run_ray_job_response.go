package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunRayJobResponse Response Object
type RunRayJobResponse struct {

	// **参数解释**：作业ID。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunRayJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunRayJobResponse struct{}"
	}

	return strings.Join([]string{"RunRayJobResponse", string(data)}, " ")
}

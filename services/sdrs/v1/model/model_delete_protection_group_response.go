package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtectionGroupResponse Response Object
type DeleteProtectionGroupResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProtectionGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectionGroupResponse", string(data)}, " ")
}

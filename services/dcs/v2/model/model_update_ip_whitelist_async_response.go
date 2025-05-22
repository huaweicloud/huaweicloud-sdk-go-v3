package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpWhitelistAsyncResponse Response Object
type UpdateIpWhitelistAsyncResponse struct {

	// **参数解释**： 设置IP白名单任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIpWhitelistAsyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpWhitelistAsyncResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpWhitelistAsyncResponse", string(data)}, " ")
}

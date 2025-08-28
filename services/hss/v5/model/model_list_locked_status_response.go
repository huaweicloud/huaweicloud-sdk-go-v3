package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLockedStatusResponse Response Object
type ListLockedStatusResponse struct {

	// **参数解释**: 主机安全防护配额资源的锁定状态。 **取值范围**: - true：已锁定，不可将按需计费的防护配额转为包年/包月。 - false：未锁定，可将按需计费的防护配额转为包年/包月。
	LockedStatus   *bool `json:"locked_status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListLockedStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLockedStatusResponse struct{}"
	}

	return strings.Join([]string{"ListLockedStatusResponse", string(data)}, " ")
}

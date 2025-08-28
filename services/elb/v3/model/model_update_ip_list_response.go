package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpListResponse Response Object
type UpdateIpListResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIpListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpListResponse", string(data)}, " ")
}

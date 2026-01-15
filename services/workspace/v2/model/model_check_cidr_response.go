package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCidrResponse Response Object
type CheckCidrResponse struct {

	// 规划的冲突网段
	ManageCidrs *[]string `json:"manage_cidrs,omitempty"`

	// 租户网段与冲突网段的重叠部分
	ConflictCidrs  *[]string `json:"conflict_cidrs,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CheckCidrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCidrResponse struct{}"
	}

	return strings.Join([]string{"CheckCidrResponse", string(data)}, " ")
}

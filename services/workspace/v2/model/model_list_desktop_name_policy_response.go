package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopNamePolicyResponse Response Object
type ListDesktopNamePolicyResponse struct {

	// 桌面名称策略信息。
	DesktopNamePolicyInfos *[]DesktopNamePolicyInfo `json:"desktop_name_policy_infos,omitempty"`

	// 查询返回总条数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopNamePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopNamePolicyResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopNamePolicyResponse", string(data)}, " ")
}

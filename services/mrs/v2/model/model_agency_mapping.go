package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户（组）与IAM委托的映射关系结构体
type AgencyMapping struct {

	// 该映射绑定的IAM委托名称。
	Agency string `json:"agency"`

	// 委托类型，分为“User”和“Group”两种。 - User表示该映射关系为针对用户的映射，identifiers中填写用户名称列表。 - Group表示该映射关系为针对用户组的映射，identifiers中填写用户组名称列表。
	IdentifierType string `json:"identifier_type"`

	// IAM委托映射的用户（组）名称列表。请前往IAM，单击“用户（组）”按钮，获取用户（组）名称列表。
	Identifiers []string `json:"identifiers"`

	// 该映射关系绑定的委托的的唯一标识码。请前往IAM，单击“委托”按钮，进入委托页面，将鼠标放置委托名称上，在弹窗中获取委托唯一标识码。
	AgencyId string `json:"agency_id"`
}

func (o AgencyMapping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyMapping struct{}"
	}

	return strings.Join([]string{"AgencyMapping", string(data)}, " ")
}

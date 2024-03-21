package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecuritySecrecyLevelsResponse Response Object
type ListSecuritySecrecyLevelsResponse struct {

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 密级列表
	Content        *[]SecrecyLevel `json:"content,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListSecuritySecrecyLevelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecuritySecrecyLevelsResponse struct{}"
	}

	return strings.Join([]string{"ListSecuritySecrecyLevelsResponse", string(data)}, " ")
}

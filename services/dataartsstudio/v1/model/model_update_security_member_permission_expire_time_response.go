package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityMemberPermissionExpireTimeResponse Response Object
type UpdateSecurityMemberPermissionExpireTimeResponse struct {

	// 更新有效期结果响应
	Results        *[]UpdatePermissionExpireTimeResultDtoResults `json:"results,omitempty"`
	HttpStatusCode int                                           `json:"-"`
}

func (o UpdateSecurityMemberPermissionExpireTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityMemberPermissionExpireTimeResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityMemberPermissionExpireTimeResponse", string(data)}, " ")
}

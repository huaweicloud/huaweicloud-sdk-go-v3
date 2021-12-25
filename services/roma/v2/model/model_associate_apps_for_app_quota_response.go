package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociateAppsForAppQuotaResponse struct {
	// 客户端应用与客户端配额绑定列表

	Applies        *[]AppQuotaAppBinding `json:"applies,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o AssociateAppsForAppQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAppsForAppQuotaResponse struct{}"
	}

	return strings.Join([]string{"AssociateAppsForAppQuotaResponse", string(data)}, " ")
}

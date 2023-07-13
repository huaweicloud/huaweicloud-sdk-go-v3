package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthorizationsResponse Response Object
type ListAuthorizationsResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 授权列表
	IncidentAuthList *[]IncidentOrderAuthV2 `json:"incident_auth_list,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o ListAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorizationsResponse", string(data)}, " ")
}

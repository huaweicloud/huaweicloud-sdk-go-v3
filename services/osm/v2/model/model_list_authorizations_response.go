package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAuthorizationsResponse struct {
	// 总数

	Total *int32 `json:"total,omitempty"`
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

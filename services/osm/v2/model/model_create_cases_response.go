package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCasesResponse struct {

	// 工单id
	IncidentId     *string `json:"incident_id,omitempty" xml:"incident_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCasesResponse struct{}"
	}

	return strings.Join([]string{"CreateCasesResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociateElbResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateElbResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateElbResponse struct{}"
	}

	return strings.Join([]string{"AssociateElbResponse", string(data)}, " ")
}

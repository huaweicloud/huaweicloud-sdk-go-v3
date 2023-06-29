package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStudyJobResponse Response Object
type CreateStudyJobResponse struct {

	// study作业id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStudyJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStudyJobResponse struct{}"
	}

	return strings.Join([]string{"CreateStudyJobResponse", string(data)}, " ")
}

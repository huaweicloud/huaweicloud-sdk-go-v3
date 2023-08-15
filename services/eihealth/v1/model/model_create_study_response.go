package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStudyResponse Response Object
type CreateStudyResponse struct {

	// study id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStudyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStudyResponse struct{}"
	}

	return strings.Join([]string{"CreateStudyResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteStudyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStudyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStudyResponse struct{}"
	}

	return strings.Join([]string{"DeleteStudyResponse", string(data)}, " ")
}

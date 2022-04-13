package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteResetTracksTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResetTracksTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResetTracksTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteResetTracksTaskResponse", string(data)}, " ")
}

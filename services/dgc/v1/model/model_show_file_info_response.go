package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowFileInfoResponse struct {
	Jobs *[]Job `json:"jobs,omitempty" xml:"jobs"`

	Scripts        *[]Script `json:"scripts,omitempty" xml:"scripts"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowFileInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowFileInfoResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecordUserViewVulTaskResponse Response Object
type RecordUserViewVulTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RecordUserViewVulTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordUserViewVulTaskResponse struct{}"
	}

	return strings.Join([]string{"RecordUserViewVulTaskResponse", string(data)}, " ")
}

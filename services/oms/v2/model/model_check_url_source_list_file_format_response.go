package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckUrlSourceListFileFormatResponse Response Object
type CheckUrlSourceListFileFormatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckUrlSourceListFileFormatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckUrlSourceListFileFormatResponse struct{}"
	}

	return strings.Join([]string{"CheckUrlSourceListFileFormatResponse", string(data)}, " ")
}

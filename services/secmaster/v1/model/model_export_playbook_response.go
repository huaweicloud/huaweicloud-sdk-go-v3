package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportPlaybookResponse Response Object
type ExportPlaybookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportPlaybookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportPlaybookResponse struct{}"
	}

	return strings.Join([]string{"ExportPlaybookResponse", string(data)}, " ")
}

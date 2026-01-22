package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImportStatusResponse Response Object
type ShowImportStatusResponse struct {
	Data           *ShowImportStatusId `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowImportStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImportStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowImportStatusResponse", string(data)}, " ")
}

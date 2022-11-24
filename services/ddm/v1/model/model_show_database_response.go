package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDatabaseResponse struct {
	Database       *GetDatabaseResponseBean `json:"database,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseResponse struct{}"
	}

	return strings.Join([]string{"ShowDatabaseResponse", string(data)}, " ")
}

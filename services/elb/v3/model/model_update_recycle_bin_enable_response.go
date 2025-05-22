package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecycleBinEnableResponse Response Object
type UpdateRecycleBinEnableResponse struct {
	RecycleBin     *RecycleBinResponseBody `json:"recycle_bin,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateRecycleBinEnableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinEnableResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinEnableResponse", string(data)}, " ")
}

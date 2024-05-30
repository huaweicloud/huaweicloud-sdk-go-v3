package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportResultResponse Response Object
type ImportResultResponse struct {
	Data           *DsProcessImportResultVo `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ImportResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportResultResponse struct{}"
	}

	return strings.Join([]string{"ImportResultResponse", string(data)}, " ")
}

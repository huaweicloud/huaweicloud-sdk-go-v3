package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveCredentialForBatchInspectionResponse Response Object
type SaveCredentialForBatchInspectionResponse struct {

	// 保存AK/SK是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SaveCredentialForBatchInspectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveCredentialForBatchInspectionResponse struct{}"
	}

	return strings.Join([]string{"SaveCredentialForBatchInspectionResponse", string(data)}, " ")
}

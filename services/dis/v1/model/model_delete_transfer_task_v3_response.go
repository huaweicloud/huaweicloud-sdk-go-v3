package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTransferTaskV3Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTransferTaskV3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransferTaskV3Response struct{}"
	}

	return strings.Join([]string{"DeleteTransferTaskV3Response", string(data)}, " ")
}

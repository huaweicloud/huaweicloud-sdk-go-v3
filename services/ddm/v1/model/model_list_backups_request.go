package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupsRequest Request Object
type ListBackupsRequest struct {
}

func (o ListBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsRequest struct{}"
	}

	return strings.Join([]string{"ListBackupsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CleanNextflowCacheRequest Request Object
type CleanNextflowCacheRequest struct {
}

func (o CleanNextflowCacheRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CleanNextflowCacheRequest struct{}"
	}

	return strings.Join([]string{"CleanNextflowCacheRequest", string(data)}, " ")
}

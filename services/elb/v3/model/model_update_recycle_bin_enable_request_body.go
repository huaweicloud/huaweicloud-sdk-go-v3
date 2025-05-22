package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecycleBinEnableRequestBody 更新回收站启用开关的请求体。
type UpdateRecycleBinEnableRequestBody struct {
	RecycleBin *RecycleBinRequestBody `json:"recycle_bin,omitempty"`
}

func (o UpdateRecycleBinEnableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinEnableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinEnableRequestBody", string(data)}, " ")
}

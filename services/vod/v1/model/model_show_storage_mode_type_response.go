package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStorageModeTypeResponse Response Object
type ShowStorageModeTypeResponse struct {

	// 降冷模式。 取值如下： - WHOLE：整个媒资粒度。 - ORIGIN：原文件粒度。
	StorageModeType *string `json:"storage_mode_type,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowStorageModeTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStorageModeTypeResponse struct{}"
	}

	return strings.Join([]string{"ShowStorageModeTypeResponse", string(data)}, " ")
}

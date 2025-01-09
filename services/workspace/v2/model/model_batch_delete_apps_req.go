package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppsReq 批量删除应用。
type BatchDeleteAppsReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 50]。
	Items []string `json:"items"`

	// 是否保留OBS桶文件，如果不保留删除应用时，会同时删除OBS桶文件(默认false)。 * 'true' - 保留OBS桶文件,不做任何操作。 * 'false' - 不保留OBS桶文件,删除应用时同时删除OBS桶文件。
	ReserveObsFile *bool `json:"reserve_obs_file,omitempty"`
}

func (o BatchDeleteAppsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppsReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppsReq", string(data)}, " ")
}

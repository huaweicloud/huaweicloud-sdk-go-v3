package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppRequest Request Object
type DeleteAppRequest struct {

	// 应用ID。
	AppId string `json:"app_id"`

	// 删除应用时是否保留OBS桶文件(默认false) * 'true' - 保留OBS桶文件,仅删除应用 * 'false' - 不保留OBS桶文件,删除应用同时删除OBS桶文件
	ReserveObsFile *string `json:"reserve_obs_file,omitempty"`
}

func (o DeleteAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlVolumeResp struct {

	// 磁盘大小。单位GB。  取值范围：10~128000，为10的整数倍。
	Size int32 `json:"size"`
}

func (o MysqlVolumeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlVolumeResp struct{}"
	}

	return strings.Join([]string{"MysqlVolumeResp", string(data)}, " ")
}

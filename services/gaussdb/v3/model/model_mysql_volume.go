package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlVolume struct {

	// 磁盘大小。默认值为10，单位GB。 取值范围：10~128000，必须为10的整数倍。
	Size string `json:"size"`
}

func (o MysqlVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlVolume struct{}"
	}

	return strings.Join([]string{"MysqlVolume", string(data)}, " ")
}

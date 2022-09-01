package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type DataSource struct {

	//   数据源类型。取值为OBS，且当前只支持OBS。
	Type string `json:"type" xml:"type"`

	Parameters *Parameters `json:"parameters" xml:"parameters"`
}

func (o DataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSource struct{}"
	}

	return strings.Join([]string{"DataSource", string(data)}, " ")
}

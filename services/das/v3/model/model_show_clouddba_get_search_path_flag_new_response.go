package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClouddbaGetSearchPathFlagNewResponse Response Object
type ShowClouddbaGetSearchPathFlagNewResponse struct {

	// 搜索路径标志，true表示开启，false表示关闭
	SearchPathFlag *bool `json:"search_path_flag,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowClouddbaGetSearchPathFlagNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClouddbaGetSearchPathFlagNewResponse struct{}"
	}

	return strings.Join([]string{"ShowClouddbaGetSearchPathFlagNewResponse", string(data)}, " ")
}

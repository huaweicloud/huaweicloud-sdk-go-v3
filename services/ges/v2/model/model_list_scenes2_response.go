package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScenes2Response Response Object
type ListScenes2Response struct {

	// scene场景分析插件信息。
	Results        *[]ListScenesRespResults `json:"results,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListScenes2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScenes2Response struct{}"
	}

	return strings.Join([]string{"ListScenes2Response", string(data)}, " ")
}

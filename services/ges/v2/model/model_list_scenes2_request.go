package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScenes2Request Request Object
type ListScenes2Request struct {

	// 场景名称。 当有且只有scene_name有值时，返回对应scene_name下的所有application详情。 当有且只有scene_name、application_name有值时，返回与application_name对应的application详情。 当scene_name、application_name、graph_id均无值时,返回所有SceneApplication
	SceneName *string `json:"scene_name,omitempty"`

	// 应用程序名字。 当有且只有scene_name、application_name有值时，返回与application_name对应的application详情。 当scene_name、application_name、graph_id均无值时,返回所有SceneApplication。
	ApplicationName *string `json:"application_name,omitempty"`

	// 图ID。 当有且只有graph_id有值时，返回对应图id下所订阅的application详情。 当scene_name、application_name、graph_id均无值时,返回所有SceneApplication。
	GraphId *string `json:"graph_id,omitempty"`
}

func (o ListScenes2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScenes2Request struct{}"
	}

	return strings.Join([]string{"ListScenes2Request", string(data)}, " ")
}

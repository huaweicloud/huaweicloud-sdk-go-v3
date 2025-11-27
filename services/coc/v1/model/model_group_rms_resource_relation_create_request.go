package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupRmsResourceRelationCreateRequest **参数解释：** 分组的关联配置信息，比如对应的APM的配置信息。 **取值范围：** 不涉及。
type GroupRmsResourceRelationCreateRequest struct {

	// **参数解释：** CloudCMDB分配的分组uuid。 **取值范围：** 不涉及。
	GroupId string `json:"group_id"`

	// **参数解释：** CloudCMDB分配的RMS资源uuid列表。 **取值范围：** 不涉及。
	CmdbResourceIdList []string `json:"cmdb_resource_id_list"`
}

func (o GroupRmsResourceRelationCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupRmsResourceRelationCreateRequest struct{}"
	}

	return strings.Join([]string{"GroupRmsResourceRelationCreateRequest", string(data)}, " ")
}

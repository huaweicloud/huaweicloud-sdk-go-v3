package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectedWdrSnapshotInfoResultObsBucket **参数解释**： 存储WDR报告临时文件的OBS桶信息。
type CollectedWdrSnapshotInfoResultObsBucket struct {

	// **参数解释**： OBS桶名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： OBS桶类型。 **取值范围**： - common：公共桶。 - aps：智能运维专用桶。
	Type *string `json:"type,omitempty"`

	// **参数解释**： OBS服务访问地址。 **取值范围**： 不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释**： OBS服务端口号。 **取值范围**： 不涉及。
	Port *string `json:"port,omitempty"`

	// **参数解释**： 最终租户ID。 **取值范围**： 不涉及。
	DomainId *string `json:"domain_id,omitempty"`
}

func (o CollectedWdrSnapshotInfoResultObsBucket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectedWdrSnapshotInfoResultObsBucket struct{}"
	}

	return strings.Join([]string{"CollectedWdrSnapshotInfoResultObsBucket", string(data)}, " ")
}

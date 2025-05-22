package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSyncRecordsResponse Response Object
type ListSyncRecordsResponse struct {

	// **参数解释**： 同步记录条数。 **取值范围**： 大于等于0的整数。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 同步记录详细信息。 **取值范围**： 不涉及。
	SyncRecords    *[]IdentitySourceSyncRecordVo `json:"sync_records,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListSyncRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyncRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListSyncRecordsResponse", string(data)}, " ")
}

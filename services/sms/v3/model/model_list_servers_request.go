package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListServersRequest struct {
	// 源端服务器状态

	State *ListServersRequestState `json:"state,omitempty"`
	// 源端服务器名称

	Name *string `json:"name,omitempty"`
	// 源端服务器ID

	Id *string `json:"id,omitempty"`
	// 源端服务器IP地址

	Ip *string `json:"ip,omitempty"`
	// 迁移项目id，填写该参数将查询迁移项目下的所有虚拟机

	Migproject *string `json:"migproject,omitempty"`
	// 每一页记录的源端服务器数量，0表示用默认值 200

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，默认值0

	Offset *int32 `json:"offset,omitempty"`
	// 根据迁移周期查询

	MigrationCycle *string `json:"migration_cycle,omitempty"`
	// 查询失去连接的源端

	Connected *bool `json:"connected,omitempty"`
	// 需要查询的企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListServersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListServersRequest struct{}"
	}

	return strings.Join([]string{"ListServersRequest", string(data)}, " ")
}

type ListServersRequestState struct {
	value string
}

type ListServersRequestStateEnum struct {
	UNAVAILABLE ListServersRequestState
	WAITING     ListServersRequestState
	INITIALIZE  ListServersRequestState
	REPLICATE   ListServersRequestState
	SYNCING     ListServersRequestState
	STOPPING    ListServersRequestState
	STOPPED     ListServersRequestState
	DELETING    ListServersRequestState
	ERROR       ListServersRequestState
	CLONING     ListServersRequestState
	CUTOVERING  ListServersRequestState
	FINISHED    ListServersRequestState
}

func GetListServersRequestStateEnum() ListServersRequestStateEnum {
	return ListServersRequestStateEnum{
		UNAVAILABLE: ListServersRequestState{
			value: "unavailable:环境校验不通过",
		},
		WAITING: ListServersRequestState{
			value: " waiting:等待",
		},
		INITIALIZE: ListServersRequestState{
			value: " initialize:初始化",
		},
		REPLICATE: ListServersRequestState{
			value: " replicate:复制",
		},
		SYNCING: ListServersRequestState{
			value: " syncing:持续同步",
		},
		STOPPING: ListServersRequestState{
			value: " stopping:暂停中",
		},
		STOPPED: ListServersRequestState{
			value: " stopped:已暂停",
		},
		DELETING: ListServersRequestState{
			value: " deleting:删除中",
		},
		ERROR: ListServersRequestState{
			value: " error:错误",
		},
		CLONING: ListServersRequestState{
			value: " cloning:等待克隆完成",
		},
		CUTOVERING: ListServersRequestState{
			value: " cutovering:启动目的端中",
		},
		FINISHED: ListServersRequestState{
			value: " finished:启动目的端完成",
		},
	}
}

func (c ListServersRequestState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListServersRequestState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

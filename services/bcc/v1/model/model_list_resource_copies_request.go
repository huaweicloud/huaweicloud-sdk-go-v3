package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceCopiesRequest Request Object
type ListResourceCopiesRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 备份类型:backup-备份, replication-复制
	CopyType *ListResourceCopiesRequestCopyType `json:"copy_type,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 备份ID
	CopyId *string `json:"copy_id,omitempty"`

	// CBR存储库ID
	VaultId *string `json:"vault_id,omitempty"`

	// 返回的最大数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页。
	Marker *string `json:"marker,omitempty"`

	// 查询开始时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListResourceCopiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceCopiesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceCopiesRequest", string(data)}, " ")
}

type ListResourceCopiesRequestCopyType struct {
	value string
}

type ListResourceCopiesRequestCopyTypeEnum struct {
	BACKUP      ListResourceCopiesRequestCopyType
	REPLICATION ListResourceCopiesRequestCopyType
}

func GetListResourceCopiesRequestCopyTypeEnum() ListResourceCopiesRequestCopyTypeEnum {
	return ListResourceCopiesRequestCopyTypeEnum{
		BACKUP: ListResourceCopiesRequestCopyType{
			value: "backup",
		},
		REPLICATION: ListResourceCopiesRequestCopyType{
			value: "replication",
		},
	}
}

func (c ListResourceCopiesRequestCopyType) Value() string {
	return c.value
}

func (c ListResourceCopiesRequestCopyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceCopiesRequestCopyType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

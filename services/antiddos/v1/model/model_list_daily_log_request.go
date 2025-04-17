package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDailyLogRequest Request Object
type ListDailyLogRequest struct {

	// 用户EIP对应的ID
	FloatingIpId string `json:"floating_ip_id"`

	// 可选范围： - desc：表示时间降序 - asc：表示时间升序 默认值为“desc”。
	SortDir *ListDailyLogRequestSortDir `json:"sort_dir,omitempty"`

	// 返回结果个数限制，此次查询返回数量最大值，取值范围：1～100，与offset配合使用。 若“limit”与“offset”均不携带则返回所有主机列表。
	Limit *string `json:"limit,omitempty"`

	// 偏移量，“limit”携带时此字段有效。
	Offset *string `json:"offset,omitempty"`

	// 用户EIP
	Ip *string `json:"ip,omitempty"`
}

func (o ListDailyLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDailyLogRequest struct{}"
	}

	return strings.Join([]string{"ListDailyLogRequest", string(data)}, " ")
}

type ListDailyLogRequestSortDir struct {
	value string
}

type ListDailyLogRequestSortDirEnum struct {
	DESC ListDailyLogRequestSortDir
	ASC  ListDailyLogRequestSortDir
}

func GetListDailyLogRequestSortDirEnum() ListDailyLogRequestSortDirEnum {
	return ListDailyLogRequestSortDirEnum{
		DESC: ListDailyLogRequestSortDir{
			value: "desc",
		},
		ASC: ListDailyLogRequestSortDir{
			value: "asc",
		},
	}
}

func (c ListDailyLogRequestSortDir) Value() string {
	return c.value
}

func (c ListDailyLogRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDailyLogRequestSortDir) UnmarshalJSON(b []byte) error {
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

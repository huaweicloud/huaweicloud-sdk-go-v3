package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListVifPeerDetectionsRequest Request Object
type ListVifPeerDetectionsRequest struct {

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListVifPeerDetectionsRequestSortDir `json:"sort_dir,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 虚拟接口对等体ID
	VifPeerId string `json:"vif_peer_id"`
}

func (o ListVifPeerDetectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVifPeerDetectionsRequest struct{}"
	}

	return strings.Join([]string{"ListVifPeerDetectionsRequest", string(data)}, " ")
}

type ListVifPeerDetectionsRequestSortDir struct {
	value string
}

type ListVifPeerDetectionsRequestSortDirEnum struct {
	ASC  ListVifPeerDetectionsRequestSortDir
	DESC ListVifPeerDetectionsRequestSortDir
}

func GetListVifPeerDetectionsRequestSortDirEnum() ListVifPeerDetectionsRequestSortDirEnum {
	return ListVifPeerDetectionsRequestSortDirEnum{
		ASC: ListVifPeerDetectionsRequestSortDir{
			value: "asc",
		},
		DESC: ListVifPeerDetectionsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListVifPeerDetectionsRequestSortDir) Value() string {
	return c.value
}

func (c ListVifPeerDetectionsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVifPeerDetectionsRequestSortDir) UnmarshalJSON(b []byte) error {
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

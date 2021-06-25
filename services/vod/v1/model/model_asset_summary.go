package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 媒资总览
type AssetSummary struct {
	// 媒资ID。

	AssetId string `json:"asset_id"`
	// 媒体标题。 长度不超过128 个字节，utf-8 编码。

	Title string `json:"title"`
	// 视频描述。 长度不超过 1024个字节。

	Description *string `json:"description,omitempty"`
	// 视频时长。单位为秒

	Duration int32 `json:"duration"`
	// 视频大小。单位为字节。

	Size int64 `json:"size"`
	// 原始播放url。

	OriginalUrl *string `json:"original_url,omitempty"`
	// 媒资分类名称。

	Category *string `json:"category,omitempty"`
	// 封面信息。

	Covers *[]CoverInfo `json:"covers,omitempty"`
	// 媒资创建时 间。 格式为 yyyymmddhhm mss。必须是 与时区无关的 UTC时间。

	CreateTime *string `json:"create_time,omitempty"`
	// 媒资状态。 \"CREATING\"   //上传中 \"FAILED\"     //上传失败 \"CREATED\"  //上传成功 \"PUBLISHED\"  //已发布 \"DELETED\"  //已删除

	AssetStatus AssetSummaryAssetStatus `json:"asset_status"`
	// 转码状态 \"UN_TRANSCODE\"        //未转码 \"WAITING_TRANSCODE\"   //等待转码，排队中 \"TRANSCODING\"         //转码中 \"TRANSCODE_SUCCEED\"   //转码成功 \"TRANSCODE_FAILED\"     //转码失败

	TranscodeStatus *AssetSummaryTranscodeStatus `json:"transcode_status,omitempty"`
	// 截图状态 \"UN_THUMBNAIL\"        //未截图 \"THUMBNAILING\"    //截图中 \"THUMBNAIL_SUCCEED\"  //截图成功 \"THUMBNAIL_FAILED\"     //截图失败

	ThumbnailStatus *AssetSummaryThumbnailStatus `json:"thumbnail_status,omitempty"`
	// 内容审核状态 \"UN_REVIEW\"        //未审核 \"REVIEWING\"  //审核中 \"REVIEW_SUSPICIOUS \"  //审核不过，待人工复审 \"REVIEW_PASSED\"      //审核通过 \"REVIEW_FAILED\"      //审核任务失败 \"REVIEW_BLOCKED\"      //已屏蔽

	ReviewStatus *AssetSummaryReviewStatus `json:"review_status,omitempty"`
	// 媒资的任务执行描述汇总。

	ExecDesc *string `json:"exec_desc,omitempty"`
	// 媒资的音视频文件格式。

	MediaType *string `json:"media_type,omitempty"`
}

func (o AssetSummary) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssetSummary struct{}"
	}

	return strings.Join([]string{"AssetSummary", string(data)}, " ")
}

type AssetSummaryAssetStatus struct {
	value string
}

type AssetSummaryAssetStatusEnum struct {
	CREATING  AssetSummaryAssetStatus
	FAILED    AssetSummaryAssetStatus
	CREATED   AssetSummaryAssetStatus
	PUBLISHED AssetSummaryAssetStatus
	DELETED   AssetSummaryAssetStatus
}

func GetAssetSummaryAssetStatusEnum() AssetSummaryAssetStatusEnum {
	return AssetSummaryAssetStatusEnum{
		CREATING: AssetSummaryAssetStatus{
			value: "CREATING",
		},
		FAILED: AssetSummaryAssetStatus{
			value: "FAILED",
		},
		CREATED: AssetSummaryAssetStatus{
			value: "CREATED",
		},
		PUBLISHED: AssetSummaryAssetStatus{
			value: "PUBLISHED",
		},
		DELETED: AssetSummaryAssetStatus{
			value: "DELETED",
		},
	}
}

func (c AssetSummaryAssetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *AssetSummaryAssetStatus) UnmarshalJSON(b []byte) error {
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

type AssetSummaryTranscodeStatus struct {
	value string
}

type AssetSummaryTranscodeStatusEnum struct {
	UN_TRANSCODE      AssetSummaryTranscodeStatus
	WAITING_TRANSCODE AssetSummaryTranscodeStatus
	TRANSCODING       AssetSummaryTranscodeStatus
	TRANSCODE_SUCCEED AssetSummaryTranscodeStatus
	TRANSCODE_FAILED  AssetSummaryTranscodeStatus
}

func GetAssetSummaryTranscodeStatusEnum() AssetSummaryTranscodeStatusEnum {
	return AssetSummaryTranscodeStatusEnum{
		UN_TRANSCODE: AssetSummaryTranscodeStatus{
			value: "UN_TRANSCODE",
		},
		WAITING_TRANSCODE: AssetSummaryTranscodeStatus{
			value: "WAITING_TRANSCODE",
		},
		TRANSCODING: AssetSummaryTranscodeStatus{
			value: "TRANSCODING",
		},
		TRANSCODE_SUCCEED: AssetSummaryTranscodeStatus{
			value: "TRANSCODE_SUCCEED",
		},
		TRANSCODE_FAILED: AssetSummaryTranscodeStatus{
			value: "TRANSCODE_FAILED",
		},
	}
}

func (c AssetSummaryTranscodeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *AssetSummaryTranscodeStatus) UnmarshalJSON(b []byte) error {
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

type AssetSummaryThumbnailStatus struct {
	value string
}

type AssetSummaryThumbnailStatusEnum struct {
	UN_THUMBNAIL      AssetSummaryThumbnailStatus
	THUMBNAILING      AssetSummaryThumbnailStatus
	THUMBNAIL_SUCCEED AssetSummaryThumbnailStatus
	THUMBNAIL_FAILED  AssetSummaryThumbnailStatus
}

func GetAssetSummaryThumbnailStatusEnum() AssetSummaryThumbnailStatusEnum {
	return AssetSummaryThumbnailStatusEnum{
		UN_THUMBNAIL: AssetSummaryThumbnailStatus{
			value: "UN_THUMBNAIL",
		},
		THUMBNAILING: AssetSummaryThumbnailStatus{
			value: "THUMBNAILING",
		},
		THUMBNAIL_SUCCEED: AssetSummaryThumbnailStatus{
			value: "THUMBNAIL_SUCCEED",
		},
		THUMBNAIL_FAILED: AssetSummaryThumbnailStatus{
			value: "THUMBNAIL_FAILED",
		},
	}
}

func (c AssetSummaryThumbnailStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *AssetSummaryThumbnailStatus) UnmarshalJSON(b []byte) error {
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

type AssetSummaryReviewStatus struct {
	value string
}

type AssetSummaryReviewStatusEnum struct {
	UN_REVIEW         AssetSummaryReviewStatus
	REVIEWING         AssetSummaryReviewStatus
	REVIEW_SUSPICIOUS AssetSummaryReviewStatus
	REVIEW_PASSED     AssetSummaryReviewStatus
	REVIEW_FAILED     AssetSummaryReviewStatus
	REVIEW_BLOCKED    AssetSummaryReviewStatus
}

func GetAssetSummaryReviewStatusEnum() AssetSummaryReviewStatusEnum {
	return AssetSummaryReviewStatusEnum{
		UN_REVIEW: AssetSummaryReviewStatus{
			value: "UN_REVIEW",
		},
		REVIEWING: AssetSummaryReviewStatus{
			value: "REVIEWING",
		},
		REVIEW_SUSPICIOUS: AssetSummaryReviewStatus{
			value: "REVIEW_SUSPICIOUS",
		},
		REVIEW_PASSED: AssetSummaryReviewStatus{
			value: "REVIEW_PASSED",
		},
		REVIEW_FAILED: AssetSummaryReviewStatus{
			value: "REVIEW_FAILED",
		},
		REVIEW_BLOCKED: AssetSummaryReviewStatus{
			value: "REVIEW_BLOCKED",
		},
	}
}

func (c AssetSummaryReviewStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *AssetSummaryReviewStatus) UnmarshalJSON(b []byte) error {
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

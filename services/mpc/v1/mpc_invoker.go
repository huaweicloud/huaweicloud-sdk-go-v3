package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mpc/v1/model"
)

type CreateAnimatedGraphicsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAnimatedGraphicsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAnimatedGraphicsTaskInvoker) Invoke() (*model.CreateAnimatedGraphicsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAnimatedGraphicsTaskResponse), nil
	}
}

type DeleteAnimatedGraphicsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAnimatedGraphicsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAnimatedGraphicsTaskInvoker) Invoke() (*model.DeleteAnimatedGraphicsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAnimatedGraphicsTaskResponse), nil
	}
}

type ListAnimatedGraphicsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAnimatedGraphicsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAnimatedGraphicsTaskInvoker) Invoke() (*model.ListAnimatedGraphicsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAnimatedGraphicsTaskResponse), nil
	}
}

type CreateAgenciesTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgenciesTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgenciesTaskInvoker) Invoke() (*model.CreateAgenciesTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgenciesTaskResponse), nil
	}
}

type ListAllBucketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllBucketsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllBucketsInvoker) Invoke() (*model.ListAllBucketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllBucketsResponse), nil
	}
}

type ListAllObsObjListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllObsObjListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllObsObjListInvoker) Invoke() (*model.ListAllObsObjListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllObsObjListResponse), nil
	}
}

type ListNotifyEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotifyEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotifyEventInvoker) Invoke() (*model.ListNotifyEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotifyEventResponse), nil
	}
}

type ListNotifySmnTopicConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotifySmnTopicConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotifySmnTopicConfigInvoker) Invoke() (*model.ListNotifySmnTopicConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotifySmnTopicConfigResponse), nil
	}
}

type NotifySmnTopicConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *NotifySmnTopicConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *NotifySmnTopicConfigInvoker) Invoke() (*model.NotifySmnTopicConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotifySmnTopicConfigResponse), nil
	}
}

type ShowAgenciesTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgenciesTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgenciesTaskInvoker) Invoke() (*model.ShowAgenciesTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgenciesTaskResponse), nil
	}
}

type UpdateBucketAuthorizedInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBucketAuthorizedInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBucketAuthorizedInvoker) Invoke() (*model.UpdateBucketAuthorizedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBucketAuthorizedResponse), nil
	}
}

type CreateEditingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEditingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEditingJobInvoker) Invoke() (*model.CreateEditingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEditingJobResponse), nil
	}
}

type DeleteEditingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEditingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEditingJobInvoker) Invoke() (*model.DeleteEditingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEditingJobResponse), nil
	}
}

type ListEditingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEditingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEditingJobInvoker) Invoke() (*model.ListEditingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEditingJobResponse), nil
	}
}

type CreateExtractTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExtractTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateExtractTaskInvoker) Invoke() (*model.CreateExtractTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExtractTaskResponse), nil
	}
}

type DeleteExtractTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteExtractTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteExtractTaskInvoker) Invoke() (*model.DeleteExtractTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteExtractTaskResponse), nil
	}
}

type ListExtractTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExtractTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExtractTaskInvoker) Invoke() (*model.ListExtractTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExtractTaskResponse), nil
	}
}

type CreateMbTasksReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMbTasksReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMbTasksReportInvoker) Invoke() (*model.CreateMbTasksReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMbTasksReportResponse), nil
	}
}

type CreateMergeChannelsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMergeChannelsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMergeChannelsTaskInvoker) Invoke() (*model.CreateMergeChannelsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMergeChannelsTaskResponse), nil
	}
}

type CreateResetTracksTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResetTracksTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResetTracksTaskInvoker) Invoke() (*model.CreateResetTracksTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResetTracksTaskResponse), nil
	}
}

type DeleteMergeChannelsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMergeChannelsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMergeChannelsTaskInvoker) Invoke() (*model.DeleteMergeChannelsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMergeChannelsTaskResponse), nil
	}
}

type DeleteResetTracksTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResetTracksTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResetTracksTaskInvoker) Invoke() (*model.DeleteResetTracksTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResetTracksTaskResponse), nil
	}
}

type ListMergeChannelsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeChannelsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeChannelsTaskInvoker) Invoke() (*model.ListMergeChannelsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeChannelsTaskResponse), nil
	}
}

type ListResetTracksTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResetTracksTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResetTracksTaskInvoker) Invoke() (*model.ListResetTracksTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResetTracksTaskResponse), nil
	}
}

type CreateMediaProcessTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMediaProcessTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMediaProcessTaskInvoker) Invoke() (*model.CreateMediaProcessTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMediaProcessTaskResponse), nil
	}
}

type DeleteMediaProcessTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMediaProcessTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMediaProcessTaskInvoker) Invoke() (*model.DeleteMediaProcessTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMediaProcessTaskResponse), nil
	}
}

type ListMediaProcessTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMediaProcessTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMediaProcessTaskInvoker) Invoke() (*model.ListMediaProcessTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMediaProcessTaskResponse), nil
	}
}

type CreateMpeCallBackInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMpeCallBackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMpeCallBackInvoker) Invoke() (*model.CreateMpeCallBackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMpeCallBackResponse), nil
	}
}

type CreateQualityEnhanceTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQualityEnhanceTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateQualityEnhanceTemplateInvoker) Invoke() (*model.CreateQualityEnhanceTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQualityEnhanceTemplateResponse), nil
	}
}

type DeleteQualityEnhanceTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteQualityEnhanceTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteQualityEnhanceTemplateInvoker) Invoke() (*model.DeleteQualityEnhanceTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQualityEnhanceTemplateResponse), nil
	}
}

type ListQualityEnhanceDefaultTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQualityEnhanceDefaultTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQualityEnhanceDefaultTemplateInvoker) Invoke() (*model.ListQualityEnhanceDefaultTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQualityEnhanceDefaultTemplateResponse), nil
	}
}

type UpdateQualityEnhanceTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateQualityEnhanceTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateQualityEnhanceTemplateInvoker) Invoke() (*model.UpdateQualityEnhanceTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateQualityEnhanceTemplateResponse), nil
	}
}

type ListTranscodeDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTranscodeDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTranscodeDetailInvoker) Invoke() (*model.ListTranscodeDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTranscodeDetailResponse), nil
	}
}

type CancelRemuxTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelRemuxTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelRemuxTaskInvoker) Invoke() (*model.CancelRemuxTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelRemuxTaskResponse), nil
	}
}

type CreateRemuxTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRemuxTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRemuxTaskInvoker) Invoke() (*model.CreateRemuxTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRemuxTaskResponse), nil
	}
}

type CreateRetryRemuxTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRetryRemuxTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRetryRemuxTaskInvoker) Invoke() (*model.CreateRetryRemuxTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRetryRemuxTaskResponse), nil
	}
}

type DeleteRemuxTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRemuxTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRemuxTaskInvoker) Invoke() (*model.DeleteRemuxTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRemuxTaskResponse), nil
	}
}

type ListRemuxTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRemuxTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRemuxTaskInvoker) Invoke() (*model.ListRemuxTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRemuxTaskResponse), nil
	}
}

type CreateTemplateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTemplateGroupInvoker) Invoke() (*model.CreateTemplateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplateGroupResponse), nil
	}
}

type DeleteTemplateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTemplateGroupInvoker) Invoke() (*model.DeleteTemplateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateGroupResponse), nil
	}
}

type ListTemplateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplateGroupInvoker) Invoke() (*model.ListTemplateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplateGroupResponse), nil
	}
}

type UpdateTemplateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTemplateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTemplateGroupInvoker) Invoke() (*model.UpdateTemplateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTemplateGroupResponse), nil
	}
}

type ShowTenantAccessInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTenantAccessInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTenantAccessInfoInvoker) Invoke() (*model.ShowTenantAccessInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTenantAccessInfoResponse), nil
	}
}

type UpdateTenantAccessInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTenantAccessInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTenantAccessInfoInvoker) Invoke() (*model.UpdateTenantAccessInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTenantAccessInfoResponse), nil
	}
}

type CreateThumbnailsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateThumbnailsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateThumbnailsTaskInvoker) Invoke() (*model.CreateThumbnailsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateThumbnailsTaskResponse), nil
	}
}

type DeleteThumbnailsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteThumbnailsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteThumbnailsTaskInvoker) Invoke() (*model.DeleteThumbnailsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteThumbnailsTaskResponse), nil
	}
}

type ListThumbnailsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListThumbnailsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListThumbnailsTaskInvoker) Invoke() (*model.ListThumbnailsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListThumbnailsTaskResponse), nil
	}
}

type CreateTranscodingTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTranscodingTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTranscodingTaskInvoker) Invoke() (*model.CreateTranscodingTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTranscodingTaskResponse), nil
	}
}

type DeleteTranscodingTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTranscodingTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTranscodingTaskInvoker) Invoke() (*model.DeleteTranscodingTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTranscodingTaskResponse), nil
	}
}

type DeleteTranscodingTaskByConsoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTranscodingTaskByConsoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTranscodingTaskByConsoleInvoker) Invoke() (*model.DeleteTranscodingTaskByConsoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTranscodingTaskByConsoleResponse), nil
	}
}

type ListStatSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStatSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStatSummaryInvoker) Invoke() (*model.ListStatSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStatSummaryResponse), nil
	}
}

type ListTranscodingTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTranscodingTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTranscodingTaskInvoker) Invoke() (*model.ListTranscodingTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTranscodingTaskResponse), nil
	}
}

type CreateTransTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTransTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTransTemplateInvoker) Invoke() (*model.CreateTransTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTransTemplateResponse), nil
	}
}

type DeleteTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTemplateInvoker) Invoke() (*model.DeleteTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateResponse), nil
	}
}

type ListTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplateInvoker) Invoke() (*model.ListTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplateResponse), nil
	}
}

type UpdateTransTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTransTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTransTemplateInvoker) Invoke() (*model.UpdateTransTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTransTemplateResponse), nil
	}
}

type CreateWatermarkTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWatermarkTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWatermarkTemplateInvoker) Invoke() (*model.CreateWatermarkTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWatermarkTemplateResponse), nil
	}
}

type DeleteWatermarkTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWatermarkTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWatermarkTemplateInvoker) Invoke() (*model.DeleteWatermarkTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWatermarkTemplateResponse), nil
	}
}

type ListWatermarkTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWatermarkTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWatermarkTemplateInvoker) Invoke() (*model.ListWatermarkTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWatermarkTemplateResponse), nil
	}
}

type UpdateWatermarkTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWatermarkTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWatermarkTemplateInvoker) Invoke() (*model.UpdateWatermarkTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWatermarkTemplateResponse), nil
	}
}

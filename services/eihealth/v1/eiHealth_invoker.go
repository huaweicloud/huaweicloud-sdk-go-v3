package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eihealth/v1/model"
)

type AddDrugDatabaseFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDrugDatabaseFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDrugDatabaseFileInvoker) Invoke() (*model.AddDrugDatabaseFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDrugDatabaseFileResponse), nil
	}
}

type BatchCancelJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCancelJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCancelJobInvoker) Invoke() (*model.BatchCancelJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCancelJobResponse), nil
	}
}

type BatchDeleteDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDataInvoker) Invoke() (*model.BatchDeleteDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDataResponse), nil
	}
}

type BatchDeleteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteJobInvoker) Invoke() (*model.BatchDeleteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteJobResponse), nil
	}
}

type BatchDeleteLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteLabelInvoker) Invoke() (*model.BatchDeleteLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteLabelResponse), nil
	}
}

type BatchDeleteMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteMemberInvoker) Invoke() (*model.BatchDeleteMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteMemberResponse), nil
	}
}

type BatchDeleteNoticeInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteNoticeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteNoticeInvoker) Invoke() (*model.BatchDeleteNoticeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteNoticeResponse), nil
	}
}

type BatchDeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteTagInvoker) Invoke() (*model.BatchDeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTagResponse), nil
	}
}

type BatchDownloadResourceStatDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDownloadResourceStatDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDownloadResourceStatDataInvoker) Invoke() (*model.BatchDownloadResourceStatDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDownloadResourceStatDataResponse), nil
	}
}

type BatchImportAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchImportAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchImportAppInvoker) Invoke() (*model.BatchImportAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchImportAppResponse), nil
	}
}

type BatchRetryJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRetryJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRetryJobInvoker) Invoke() (*model.BatchRetryJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRetryJobResponse), nil
	}
}

type BatchUpdateNodeLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateNodeLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateNodeLabelInvoker) Invoke() (*model.BatchUpdateNodeLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateNodeLabelResponse), nil
	}
}

type BatchUpdateNoticeInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateNoticeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateNoticeInvoker) Invoke() (*model.BatchUpdateNoticeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateNoticeResponse), nil
	}
}

type CancelDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelDataJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelDataJobInvoker) Invoke() (*model.CancelDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelDataJobResponse), nil
	}
}

type CancelDrugJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelDrugJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelDrugJobInvoker) Invoke() (*model.CancelDrugJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelDrugJobResponse), nil
	}
}

type CancelJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelJobInvoker) Invoke() (*model.CancelJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelJobResponse), nil
	}
}

type ChangePasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePasswordInvoker) Invoke() (*model.ChangePasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePasswordResponse), nil
	}
}

type CheckEmailConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckEmailConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckEmailConnectionInvoker) Invoke() (*model.CheckEmailConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckEmailConnectionResponse), nil
	}
}

type CheckTokenVerificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckTokenVerificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckTokenVerificationInvoker) Invoke() (*model.CheckTokenVerificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckTokenVerificationResponse), nil
	}
}

type CopyDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyDataInvoker) Invoke() (*model.CopyDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyDataResponse), nil
	}
}

type CreateAdmetJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAdmetJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAdmetJobInvoker) Invoke() (*model.CreateAdmetJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAdmetJobResponse), nil
	}
}

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

type CreateAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAutoJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAutoJobInvoker) Invoke() (*model.CreateAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAutoJobResponse), nil
	}
}

type CreateBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBackupInvoker) Invoke() (*model.CreateBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBackupResponse), nil
	}
}

type CreateClusterJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterJobInvoker) Invoke() (*model.CreateClusterJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterJobResponse), nil
	}
}

type CreateCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCodeInvoker) Invoke() (*model.CreateCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCodeResponse), nil
	}
}

type CreateComputingResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateComputingResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateComputingResourceInvoker) Invoke() (*model.CreateComputingResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateComputingResourceResponse), nil
	}
}

type CreateDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataInvoker) Invoke() (*model.CreateDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataResponse), nil
	}
}

type CreateDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDatabaseDataInvoker) Invoke() (*model.CreateDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseDataResponse), nil
	}
}

type CreateDatabaseResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDatabaseResourceInvoker) Invoke() (*model.CreateDatabaseResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseResourceResponse), nil
	}
}

type CreateDockingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDockingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDockingJobInvoker) Invoke() (*model.CreateDockingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDockingJobResponse), nil
	}
}

type CreateDrugDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDrugDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDrugDatabaseInvoker) Invoke() (*model.CreateDrugDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDrugDatabaseResponse), nil
	}
}

type CreateDrugModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDrugModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDrugModelInvoker) Invoke() (*model.CreateDrugModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDrugModelResponse), nil
	}
}

type CreateFepJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFepJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFepJobInvoker) Invoke() (*model.CreateFepJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFepJobResponse), nil
	}
}

type CreateGenJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGenJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGenJobInvoker) Invoke() (*model.CreateGenJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGenJobResponse), nil
	}
}

type CreateImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateImageInvoker) Invoke() (*model.CreateImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type CreateLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLabelInvoker) Invoke() (*model.CreateLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLabelResponse), nil
	}
}

type CreateLabelPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLabelPageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLabelPageInvoker) Invoke() (*model.CreateLabelPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLabelPageResponse), nil
	}
}

type CreateMolBatchDownloadTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMolBatchDownloadTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMolBatchDownloadTaskInvoker) Invoke() (*model.CreateMolBatchDownloadTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMolBatchDownloadTaskResponse), nil
	}
}

type CreateOptmJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOptmJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOptmJobInvoker) Invoke() (*model.CreateOptmJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOptmJobResponse), nil
	}
}

type CreatePerformanceResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePerformanceResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePerformanceResourceInvoker) Invoke() (*model.CreatePerformanceResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePerformanceResourceResponse), nil
	}
}

type CreatePocketDetectionJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePocketDetectionJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePocketDetectionJobInvoker) Invoke() (*model.CreatePocketDetectionJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePocketDetectionJobResponse), nil
	}
}

type CreatePocketMolDesignJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePocketMolDesignJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePocketMolDesignJobInvoker) Invoke() (*model.CreatePocketMolDesignJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePocketMolDesignJobResponse), nil
	}
}

type CreateProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectInvoker) Invoke() (*model.CreateProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectResponse), nil
	}
}

type CreateScaleOutPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScaleOutPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScaleOutPolicyInvoker) Invoke() (*model.CreateScaleOutPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScaleOutPolicyResponse), nil
	}
}

type CreateSearchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSearchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSearchJobInvoker) Invoke() (*model.CreateSearchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSearchJobResponse), nil
	}
}

type CreateStudyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStudyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStudyInvoker) Invoke() (*model.CreateStudyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStudyResponse), nil
	}
}

type CreateStudyJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStudyJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStudyJobInvoker) Invoke() (*model.CreateStudyJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStudyJobResponse), nil
	}
}

type CreateSynthesisJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSynthesisJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSynthesisJobInvoker) Invoke() (*model.CreateSynthesisJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSynthesisJobResponse), nil
	}
}

type CreateTargetOptJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTargetOptJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTargetOptJobInvoker) Invoke() (*model.CreateTargetOptJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTargetOptJobResponse), nil
	}
}

type CreateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTemplateInvoker) Invoke() (*model.CreateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplateResponse), nil
	}
}

type CreateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUserInvoker) Invoke() (*model.CreateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserResponse), nil
	}
}

type CreateWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowInvoker) Invoke() (*model.CreateWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowResponse), nil
	}
}

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
	}
}

type DeleteAssetVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAssetVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAssetVersionInvoker) Invoke() (*model.DeleteAssetVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAssetVersionResponse), nil
	}
}

type DeleteAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAutoJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAutoJobInvoker) Invoke() (*model.DeleteAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAutoJobResponse), nil
	}
}

type DeleteBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBackupInvoker) Invoke() (*model.DeleteBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackupResponse), nil
	}
}

type DeleteComputingResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteComputingResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteComputingResourceInvoker) Invoke() (*model.DeleteComputingResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteComputingResourceResponse), nil
	}
}

type DeleteDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataJobInvoker) Invoke() (*model.DeleteDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataJobResponse), nil
	}
}

type DeleteDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDatabaseDataInvoker) Invoke() (*model.DeleteDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseDataResponse), nil
	}
}

type DeleteDatabaseResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDatabaseResourceInvoker) Invoke() (*model.DeleteDatabaseResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseResourceResponse), nil
	}
}

type DeleteDrugDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDrugDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDrugDatabaseInvoker) Invoke() (*model.DeleteDrugDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDrugDatabaseResponse), nil
	}
}

type DeleteDrugJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDrugJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDrugJobInvoker) Invoke() (*model.DeleteDrugJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDrugJobResponse), nil
	}
}

type DeleteDrugModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDrugModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDrugModelInvoker) Invoke() (*model.DeleteDrugModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDrugModelResponse), nil
	}
}

type DeleteImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteImageInvoker) Invoke() (*model.DeleteImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type DeleteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteJobInvoker) Invoke() (*model.DeleteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteJobResponse), nil
	}
}

type DeleteLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLabelInvoker) Invoke() (*model.DeleteLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLabelResponse), nil
	}
}

type DeleteLabelPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLabelPageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLabelPageInvoker) Invoke() (*model.DeleteLabelPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLabelPageResponse), nil
	}
}

type DeleteMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMemberInvoker) Invoke() (*model.DeleteMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMemberResponse), nil
	}
}

type DeleteMessageEmailConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMessageEmailConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMessageEmailConfigInvoker) Invoke() (*model.DeleteMessageEmailConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMessageEmailConfigResponse), nil
	}
}

type DeletePerformanceResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePerformanceResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePerformanceResourceInvoker) Invoke() (*model.DeletePerformanceResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePerformanceResourceResponse), nil
	}
}

type DeleteProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProjectInvoker) Invoke() (*model.DeleteProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProjectResponse), nil
	}
}

type DeleteScaleOutPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScaleOutPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScaleOutPolicyInvoker) Invoke() (*model.DeleteScaleOutPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScaleOutPolicyResponse), nil
	}
}

type DeleteStarInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStarInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStarInvoker) Invoke() (*model.DeleteStarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStarResponse), nil
	}
}

type DeleteStudyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStudyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStudyInvoker) Invoke() (*model.DeleteStudyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStudyResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
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

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type DeleteWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkflowInvoker) Invoke() (*model.DeleteWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkflowResponse), nil
	}
}

type DownloadDataJobLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadDataJobLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadDataJobLogInvoker) Invoke() (*model.DownloadDataJobLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadDataJobLogResponse), nil
	}
}

type DownloadDataTraceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadDataTraceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadDataTraceInvoker) Invoke() (*model.DownloadDataTraceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadDataTraceResponse), nil
	}
}

type ExecuteAssetActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteAssetActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteAssetActionInvoker) Invoke() (*model.ExecuteAssetActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteAssetActionResponse), nil
	}
}

type ExecuteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteJobInvoker) Invoke() (*model.ExecuteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteJobResponse), nil
	}
}

type GenerateComplexCombineInvoker struct {
	*invoker.BaseInvoker
}

func (i *GenerateComplexCombineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GenerateComplexCombineInvoker) Invoke() (*model.GenerateComplexCombineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenerateComplexCombineResponse), nil
	}
}

type GeneratePocketFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *GeneratePocketFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GeneratePocketFileInvoker) Invoke() (*model.GeneratePocketFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GeneratePocketFileResponse), nil
	}
}

type GenerateSurfacePointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *GenerateSurfacePointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GenerateSurfacePointsInvoker) Invoke() (*model.GenerateSurfacePointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenerateSurfacePointsResponse), nil
	}
}

type ImportDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportDataInvoker) Invoke() (*model.ImportDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportDataResponse), nil
	}
}

type ImportDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportDatabaseDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportDatabaseDataInvoker) Invoke() (*model.ImportDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportDatabaseDataResponse), nil
	}
}

type ImportImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportImageInvoker) Invoke() (*model.ImportImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportImageResponse), nil
	}
}

type ImportNetworkDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportNetworkDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportNetworkDataInvoker) Invoke() (*model.ImportNetworkDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportNetworkDataResponse), nil
	}
}

type ImportTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportTemplateInvoker) Invoke() (*model.ImportTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportTemplateResponse), nil
	}
}

type ImportUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportUserInvoker) Invoke() (*model.ImportUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportUserResponse), nil
	}
}

type ImportWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportWorkflowInvoker) Invoke() (*model.ImportWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportWorkflowResponse), nil
	}
}

type ListAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppInvoker) Invoke() (*model.ListAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppResponse), nil
	}
}

type ListArchiveConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListArchiveConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListArchiveConfigsInvoker) Invoke() (*model.ListArchiveConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListArchiveConfigsResponse), nil
	}
}

type ListAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssetInvoker) Invoke() (*model.ListAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssetResponse), nil
	}
}

type ListAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAutoJobInvoker) Invoke() (*model.ListAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoJobResponse), nil
	}
}

type ListBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackupInvoker) Invoke() (*model.ListBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupResponse), nil
	}
}

type ListBaseModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBaseModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBaseModelInvoker) Invoke() (*model.ListBaseModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBaseModelResponse), nil
	}
}

type ListBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBucketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBucketInvoker) Invoke() (*model.ListBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBucketResponse), nil
	}
}

type ListCheckpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCheckpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCheckpointInvoker) Invoke() (*model.ListCheckpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCheckpointResponse), nil
	}
}

type ListClusterAllNodeLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterAllNodeLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterAllNodeLabelInvoker) Invoke() (*model.ListClusterAllNodeLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterAllNodeLabelResponse), nil
	}
}

type ListComputingResourceFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComputingResourceFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListComputingResourceFlavorsInvoker) Invoke() (*model.ListComputingResourceFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComputingResourceFlavorsResponse), nil
	}
}

type ListComputingResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComputingResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListComputingResourcesInvoker) Invoke() (*model.ListComputingResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComputingResourcesResponse), nil
	}
}

type ListDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataInvoker) Invoke() (*model.ListDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataResponse), nil
	}
}

type ListDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataJobInvoker) Invoke() (*model.ListDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataJobResponse), nil
	}
}

type ListDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatabaseDataInvoker) Invoke() (*model.ListDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseDataResponse), nil
	}
}

type ListDatabaseResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatabaseResourceInvoker) Invoke() (*model.ListDatabaseResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseResourceResponse), nil
	}
}

type ListDatabaseResourceFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseResourceFlavorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatabaseResourceFlavorInvoker) Invoke() (*model.ListDatabaseResourceFlavorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseResourceFlavorResponse), nil
	}
}

type ListDrugDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDrugDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDrugDatabaseInvoker) Invoke() (*model.ListDrugDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDrugDatabaseResponse), nil
	}
}

type ListDrugJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDrugJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDrugJobInvoker) Invoke() (*model.ListDrugJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDrugJobResponse), nil
	}
}

type ListDrugModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDrugModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDrugModelInvoker) Invoke() (*model.ListDrugModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDrugModelResponse), nil
	}
}

type ListGlobalWorkflowStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalWorkflowStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalWorkflowStatisticInvoker) Invoke() (*model.ListGlobalWorkflowStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalWorkflowStatisticResponse), nil
	}
}

type ListIamGroupUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIamGroupUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIamGroupUsersInvoker) Invoke() (*model.ListIamGroupUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIamGroupUsersResponse), nil
	}
}

type ListIamGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIamGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIamGroupsInvoker) Invoke() (*model.ListIamGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIamGroupsResponse), nil
	}
}

type ListIamUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIamUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIamUsersInvoker) Invoke() (*model.ListIamUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIamUsersResponse), nil
	}
}

type ListImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageInvoker) Invoke() (*model.ListImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageResponse), nil
	}
}

type ListImageTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageTagInvoker) Invoke() (*model.ListImageTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageTagResponse), nil
	}
}

type ListInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceInvoker) Invoke() (*model.ListInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceResponse), nil
	}
}

type ListJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobInvoker) Invoke() (*model.ListJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobResponse), nil
	}
}

type ListLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLabelInvoker) Invoke() (*model.ListLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLabelResponse), nil
	}
}

type ListLabelPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLabelPageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLabelPageInvoker) Invoke() (*model.ListLabelPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLabelPageResponse), nil
	}
}

type ListMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessageInvoker) Invoke() (*model.ListMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageResponse), nil
	}
}

type ListMessageStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessageStatisticsInvoker) Invoke() (*model.ListMessageStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageStatisticsResponse), nil
	}
}

type ListMfaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMfaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMfaInvoker) Invoke() (*model.ListMfaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMfaResponse), nil
	}
}

type ListNodeLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodeLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNodeLabelInvoker) Invoke() (*model.ListNodeLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodeLabelResponse), nil
	}
}

type ListNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNodesInvoker) Invoke() (*model.ListNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodesResponse), nil
	}
}

type ListNoticeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNoticeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNoticeInvoker) Invoke() (*model.ListNoticeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNoticeResponse), nil
	}
}

type ListPerformanceResourceStatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPerformanceResourceStatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPerformanceResourceStatInvoker) Invoke() (*model.ListPerformanceResourceStatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPerformanceResourceStatResponse), nil
	}
}

type ListPerformanceResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPerformanceResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPerformanceResourcesInvoker) Invoke() (*model.ListPerformanceResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPerformanceResourcesResponse), nil
	}
}

type ListPolicyEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyEventsInvoker) Invoke() (*model.ListPolicyEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyEventsResponse), nil
	}
}

type ListPresetLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPresetLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPresetLabelInvoker) Invoke() (*model.ListPresetLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPresetLabelResponse), nil
	}
}

type ListProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectInvoker) Invoke() (*model.ListProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectResponse), nil
	}
}

type ListPropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPropertyInvoker) Invoke() (*model.ListPropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPropertyResponse), nil
	}
}

type ListQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotaInvoker) Invoke() (*model.ListQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaResponse), nil
	}
}

type ListScaleOutPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScaleOutPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScaleOutPolicyInvoker) Invoke() (*model.ListScaleOutPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScaleOutPolicyResponse), nil
	}
}

type ListScalingHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScalingHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScalingHistoryInvoker) Invoke() (*model.ListScalingHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScalingHistoryResponse), nil
	}
}

type ListStarInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStarInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStarInvoker) Invoke() (*model.ListStarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStarResponse), nil
	}
}

type ListStorageResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStorageResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStorageResourcesInvoker) Invoke() (*model.ListStorageResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStorageResourcesResponse), nil
	}
}

type ListStudyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStudyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStudyInvoker) Invoke() (*model.ListStudyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStudyResponse), nil
	}
}

type ListStudyJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStudyJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStudyJobInvoker) Invoke() (*model.ListStudyJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStudyJobResponse), nil
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

type ListUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserInvoker) Invoke() (*model.ListUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserResponse), nil
	}
}

type ListVendorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVendorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVendorInvoker) Invoke() (*model.ListVendorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVendorResponse), nil
	}
}

type ListWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowInvoker) Invoke() (*model.ListWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowResponse), nil
	}
}

type ListWorkflowStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowStatisticInvoker) Invoke() (*model.ListWorkflowStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowStatisticResponse), nil
	}
}

type PublishAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishAppInvoker) Invoke() (*model.PublishAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishAppResponse), nil
	}
}

type PublishDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishDataInvoker) Invoke() (*model.PublishDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishDataResponse), nil
	}
}

type PublishImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishImageInvoker) Invoke() (*model.PublishImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishImageResponse), nil
	}
}

type PublishWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishWorkflowInvoker) Invoke() (*model.PublishWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishWorkflowResponse), nil
	}
}

type QuoteDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *QuoteDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *QuoteDataInvoker) Invoke() (*model.QuoteDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuoteDataResponse), nil
	}
}

type QuoteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *QuoteInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *QuoteInstanceInvoker) Invoke() (*model.QuoteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuoteInstanceResponse), nil
	}
}

type RebootNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebootNodeInvoker) Invoke() (*model.RebootNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootNodeResponse), nil
	}
}

type RestoreBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestoreBackupInvoker) Invoke() (*model.RestoreBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreBackupResponse), nil
	}
}

type RetryDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryDataJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryDataJobInvoker) Invoke() (*model.RetryDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryDataJobResponse), nil
	}
}

type RetryJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryJobInvoker) Invoke() (*model.RetryJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryJobResponse), nil
	}
}

type RunFastaPreprocessInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunFastaPreprocessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunFastaPreprocessInvoker) Invoke() (*model.RunFastaPreprocessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunFastaPreprocessResponse), nil
	}
}

type Show3dStructureContentInvoker struct {
	*invoker.BaseInvoker
}

func (i *Show3dStructureContentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *Show3dStructureContentInvoker) Invoke() (*model.Show3dStructureContentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Show3dStructureContentResponse), nil
	}
}

type ShowAdmetJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAdmetJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAdmetJobInvoker) Invoke() (*model.ShowAdmetJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAdmetJobResponse), nil
	}
}

type ShowAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppInvoker) Invoke() (*model.ShowAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppResponse), nil
	}
}

type ShowAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssetInvoker) Invoke() (*model.ShowAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetResponse), nil
	}
}

type ShowAssetVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssetVersionInvoker) Invoke() (*model.ShowAssetVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetVersionResponse), nil
	}
}

type ShowAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoJobInvoker) Invoke() (*model.ShowAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoJobResponse), nil
	}
}

type ShowBackupPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupPathInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBackupPathInvoker) Invoke() (*model.ShowBackupPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupPathResponse), nil
	}
}

type ShowBmsDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBmsDevicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBmsDevicesInvoker) Invoke() (*model.ShowBmsDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBmsDevicesResponse), nil
	}
}

type ShowBucketStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBucketStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBucketStorageInvoker) Invoke() (*model.ShowBucketStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBucketStorageResponse), nil
	}
}

type ShowDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataInvoker) Invoke() (*model.ShowDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataResponse), nil
	}
}

type ShowDataJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataJobInvoker) Invoke() (*model.ShowDataJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataJobResponse), nil
	}
}

type ShowDataPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataPolicyInvoker) Invoke() (*model.ShowDataPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataPolicyResponse), nil
	}
}

type ShowDockerLoginInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDockerLoginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDockerLoginInvoker) Invoke() (*model.ShowDockerLoginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDockerLoginResponse), nil
	}
}

type ShowDockingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDockingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDockingJobInvoker) Invoke() (*model.ShowDockingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDockingJobResponse), nil
	}
}

type ShowEnvInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnvInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEnvInvoker) Invoke() (*model.ShowEnvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnvResponse), nil
	}
}

type ShowEvsQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEvsQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEvsQuotaInvoker) Invoke() (*model.ShowEvsQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEvsQuotaResponse), nil
	}
}

type ShowExtremumInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExtremumInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExtremumInfoInvoker) Invoke() (*model.ShowExtremumInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExtremumInfoResponse), nil
	}
}

type ShowFepJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFepJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFepJobInvoker) Invoke() (*model.ShowFepJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFepJobResponse), nil
	}
}

type ShowGenJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGenJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGenJobInvoker) Invoke() (*model.ShowGenJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGenJobResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowJobConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobConfigInvoker) Invoke() (*model.ShowJobConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobConfigResponse), nil
	}
}

type ShowJobEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobEventInvoker) Invoke() (*model.ShowJobEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobEventResponse), nil
	}
}

type ShowJobLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobLogInvoker) Invoke() (*model.ShowJobLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobLogResponse), nil
	}
}

type ShowLeftQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLeftQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLeftQuotaInvoker) Invoke() (*model.ShowLeftQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLeftQuotaResponse), nil
	}
}

type ShowMessageClearRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMessageClearRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMessageClearRuleInvoker) Invoke() (*model.ShowMessageClearRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMessageClearRuleResponse), nil
	}
}

type ShowMessageEmailConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMessageEmailConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMessageEmailConfigInvoker) Invoke() (*model.ShowMessageEmailConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMessageEmailConfigResponse), nil
	}
}

type ShowMessageReceiveConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMessageReceiveConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMessageReceiveConfigInvoker) Invoke() (*model.ShowMessageReceiveConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMessageReceiveConfigResponse), nil
	}
}

type ShowMolBatchDownloadTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMolBatchDownloadTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMolBatchDownloadTaskInvoker) Invoke() (*model.ShowMolBatchDownloadTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMolBatchDownloadTaskResponse), nil
	}
}

type ShowOptmJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOptmJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOptmJobInvoker) Invoke() (*model.ShowOptmJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOptmJobResponse), nil
	}
}

type ShowPocketDetectionJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPocketDetectionJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPocketDetectionJobInvoker) Invoke() (*model.ShowPocketDetectionJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPocketDetectionJobResponse), nil
	}
}

type ShowPocketMolDesignJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPocketMolDesignJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPocketMolDesignJobInvoker) Invoke() (*model.ShowPocketMolDesignJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPocketMolDesignJobResponse), nil
	}
}

type ShowProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectInvoker) Invoke() (*model.ShowProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectResponse), nil
	}
}

type ShowProjectTraceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectTraceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectTraceInvoker) Invoke() (*model.ShowProjectTraceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectTraceResponse), nil
	}
}

type ShowProjectTraceDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectTraceDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectTraceDataInvoker) Invoke() (*model.ShowProjectTraceDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectTraceDataResponse), nil
	}
}

type ShowProjectTrackerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectTrackerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectTrackerInvoker) Invoke() (*model.ShowProjectTrackerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectTrackerResponse), nil
	}
}

type ShowResourceMetricDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceMetricDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceMetricDataInvoker) Invoke() (*model.ShowResourceMetricDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceMetricDataResponse), nil
	}
}

type ShowScaleInPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScaleInPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScaleInPolicyInvoker) Invoke() (*model.ShowScaleInPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScaleInPolicyResponse), nil
	}
}

type ShowScaleOutPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScaleOutPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScaleOutPolicyInvoker) Invoke() (*model.ShowScaleOutPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScaleOutPolicyResponse), nil
	}
}

type ShowScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScheduleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScheduleInvoker) Invoke() (*model.ShowScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScheduleResponse), nil
	}
}

type ShowSearchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSearchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSearchJobInvoker) Invoke() (*model.ShowSearchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSearchJobResponse), nil
	}
}

type ShowSynthesisJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSynthesisJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSynthesisJobInvoker) Invoke() (*model.ShowSynthesisJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSynthesisJobResponse), nil
	}
}

type ShowTargetOptJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTargetOptJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTargetOptJobInvoker) Invoke() (*model.ShowTargetOptJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTargetOptJobResponse), nil
	}
}

type ShowTaskEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskEventsInvoker) Invoke() (*model.ShowTaskEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskEventsResponse), nil
	}
}

type ShowTaskInstanceEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInstanceEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskInstanceEventsInvoker) Invoke() (*model.ShowTaskInstanceEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskInstanceEventsResponse), nil
	}
}

type ShowTaskInstanceMetricDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInstanceMetricDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskInstanceMetricDataInvoker) Invoke() (*model.ShowTaskInstanceMetricDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskInstanceMetricDataResponse), nil
	}
}

type ShowTaskInstancePodInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInstancePodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskInstancePodInvoker) Invoke() (*model.ShowTaskInstancePodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskInstancePodResponse), nil
	}
}

type ShowTaskInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskInstancesInvoker) Invoke() (*model.ShowTaskInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskInstancesResponse), nil
	}
}

type ShowTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTemplateInvoker) Invoke() (*model.ShowTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateResponse), nil
	}
}

type ShowUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserInvoker) Invoke() (*model.ShowUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserResponse), nil
	}
}

type ShowUserSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserSettingInvoker) Invoke() (*model.ShowUserSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserSettingResponse), nil
	}
}

type ShowVendorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVendorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVendorInvoker) Invoke() (*model.ShowVendorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVendorResponse), nil
	}
}

type ShowWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowInvoker) Invoke() (*model.ShowWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowResponse), nil
	}
}

type StartAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAutoJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartAutoJobInvoker) Invoke() (*model.StartAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAutoJobResponse), nil
	}
}

type StartNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartNodeInvoker) Invoke() (*model.StartNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartNodeResponse), nil
	}
}

type StartScaleOutPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartScaleOutPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartScaleOutPolicyInvoker) Invoke() (*model.StartScaleOutPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartScaleOutPolicyResponse), nil
	}
}

type StopAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopAutoJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopAutoJobInvoker) Invoke() (*model.StopAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopAutoJobResponse), nil
	}
}

type StopNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopNodeInvoker) Invoke() (*model.StopNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopNodeResponse), nil
	}
}

type StopScaleOutPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopScaleOutPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopScaleOutPolicyInvoker) Invoke() (*model.StopScaleOutPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopScaleOutPolicyResponse), nil
	}
}

type SubscribeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SubscribeAppInvoker) Invoke() (*model.SubscribeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeAppResponse), nil
	}
}

type SubscribeDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SubscribeDataInvoker) Invoke() (*model.SubscribeDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeDataResponse), nil
	}
}

type SubscribeImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SubscribeImageInvoker) Invoke() (*model.SubscribeImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeImageResponse), nil
	}
}

type SubscribeWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SubscribeWorkflowInvoker) Invoke() (*model.SubscribeWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeWorkflowResponse), nil
	}
}

type TransferProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *TransferProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TransferProjectInvoker) Invoke() (*model.TransferProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TransferProjectResponse), nil
	}
}

type UpdateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppInvoker) Invoke() (*model.UpdateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppResponse), nil
	}
}

type UpdateArchiveConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateArchiveConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateArchiveConfigInvoker) Invoke() (*model.UpdateArchiveConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateArchiveConfigResponse), nil
	}
}

type UpdateAssetVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssetVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAssetVersionInvoker) Invoke() (*model.UpdateAssetVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssetVersionResponse), nil
	}
}

type UpdateAutoJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAutoJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAutoJobInvoker) Invoke() (*model.UpdateAutoJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAutoJobResponse), nil
	}
}

type UpdateDataPathPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataPathPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDataPathPolicyInvoker) Invoke() (*model.UpdateDataPathPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataPathPolicyResponse), nil
	}
}

type UpdateDataPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDataPolicyInvoker) Invoke() (*model.UpdateDataPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataPolicyResponse), nil
	}
}

type UpdateDatabaseDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDatabaseDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDatabaseDataInvoker) Invoke() (*model.UpdateDatabaseDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatabaseDataResponse), nil
	}
}

type UpdateDrugDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDrugDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDrugDatabaseInvoker) Invoke() (*model.UpdateDrugDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDrugDatabaseResponse), nil
	}
}

type UpdateDrugJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDrugJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDrugJobInvoker) Invoke() (*model.UpdateDrugJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDrugJobResponse), nil
	}
}

type UpdateDrugModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDrugModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDrugModelInvoker) Invoke() (*model.UpdateDrugModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDrugModelResponse), nil
	}
}

type UpdateImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateImageInvoker) Invoke() (*model.UpdateImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateImageResponse), nil
	}
}

type UpdateInitPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInitPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInitPasswordInvoker) Invoke() (*model.UpdateInitPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInitPasswordResponse), nil
	}
}

type UpdateJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateJobInvoker) Invoke() (*model.UpdateJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateJobResponse), nil
	}
}

type UpdateJobConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateJobConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateJobConfigInvoker) Invoke() (*model.UpdateJobConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateJobConfigResponse), nil
	}
}

type UpdateMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMemberInvoker) Invoke() (*model.UpdateMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMemberResponse), nil
	}
}

type UpdateMessageClearRuleRequestBodyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMessageClearRuleRequestBodyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMessageClearRuleRequestBodyInvoker) Invoke() (*model.UpdateMessageClearRuleRequestBodyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMessageClearRuleRequestBodyResponse), nil
	}
}

type UpdateMessageEmailConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMessageEmailConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMessageEmailConfigInvoker) Invoke() (*model.UpdateMessageEmailConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMessageEmailConfigResponse), nil
	}
}

type UpdateMessageReceiveConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMessageReceiveConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMessageReceiveConfigInvoker) Invoke() (*model.UpdateMessageReceiveConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMessageReceiveConfigResponse), nil
	}
}

type UpdatePerformanceResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePerformanceResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePerformanceResourceInvoker) Invoke() (*model.UpdatePerformanceResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePerformanceResourceResponse), nil
	}
}

type UpdateProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectInvoker) Invoke() (*model.UpdateProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectResponse), nil
	}
}

type UpdateProjectTrackerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectTrackerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectTrackerInvoker) Invoke() (*model.UpdateProjectTrackerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectTrackerResponse), nil
	}
}

type UpdateScaleInPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScaleInPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScaleInPolicyInvoker) Invoke() (*model.UpdateScaleInPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScaleInPolicyResponse), nil
	}
}

type UpdateScaleOutPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScaleOutPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScaleOutPolicyInvoker) Invoke() (*model.UpdateScaleOutPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScaleOutPolicyResponse), nil
	}
}

type UpdateScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScheduleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScheduleInvoker) Invoke() (*model.UpdateScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScheduleResponse), nil
	}
}

type UpdateStarInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStarInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStarInvoker) Invoke() (*model.UpdateStarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStarResponse), nil
	}
}

type UpdateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserInvoker) Invoke() (*model.UpdateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserResponse), nil
	}
}

type UpdateUserByDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserByDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserByDomainInvoker) Invoke() (*model.UpdateUserByDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserByDomainResponse), nil
	}
}

type UpdateUserRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserRoleInvoker) Invoke() (*model.UpdateUserRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserRoleResponse), nil
	}
}

type UpdateUserSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserSettingInvoker) Invoke() (*model.UpdateUserSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserSettingResponse), nil
	}
}

type UpdateVendorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVendorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVendorInvoker) Invoke() (*model.UpdateVendorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVendorResponse), nil
	}
}

type UpdateWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkflowInvoker) Invoke() (*model.UpdateWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkflowResponse), nil
	}
}

type UploadDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadDataInvoker) Invoke() (*model.UploadDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadDataResponse), nil
	}
}

type UploadTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadTemplateInvoker) Invoke() (*model.UploadTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadTemplateResponse), nil
	}
}

type ValidateCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateCodeInvoker) Invoke() (*model.ValidateCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateCodeResponse), nil
	}
}

type ShowAdmetPropertiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAdmetPropertiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAdmetPropertiesInvoker) Invoke() (*model.ShowAdmetPropertiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAdmetPropertiesResponse), nil
	}
}

type CreateCpiJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCpiJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCpiJobInvoker) Invoke() (*model.CreateCpiJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCpiJobResponse), nil
	}
}

type CreateCpiTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCpiTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCpiTaskInvoker) Invoke() (*model.CreateCpiTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCpiTaskResponse), nil
	}
}

type ShowCpiJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCpiJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCpiJobInvoker) Invoke() (*model.ShowCpiJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCpiJobResponse), nil
	}
}

type ShowCpiTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCpiTaskResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCpiTaskResultInvoker) Invoke() (*model.ShowCpiTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCpiTaskResultResponse), nil
	}
}

type CreateCssClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCssClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCssClusterInvoker) Invoke() (*model.CreateCssClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCssClusterResponse), nil
	}
}

type DeleteCssClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCssClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCssClusterInvoker) Invoke() (*model.DeleteCssClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCssClusterResponse), nil
	}
}

type ListCssClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCssClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCssClusterInvoker) Invoke() (*model.ListCssClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCssClusterResponse), nil
	}
}

type ListTermTenantCssClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTermTenantCssClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTermTenantCssClusterInvoker) Invoke() (*model.ListTermTenantCssClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTermTenantCssClusterResponse), nil
	}
}

type ValidateCssConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateCssConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateCssConnectionInvoker) Invoke() (*model.ValidateCssConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateCssConnectionResponse), nil
	}
}

type CreateCustomPropsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCustomPropsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCustomPropsTaskInvoker) Invoke() (*model.CreateCustomPropsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCustomPropsTaskResponse), nil
	}
}

type ShowCustomPropsTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomPropsTaskResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCustomPropsTaskResultInvoker) Invoke() (*model.ShowCustomPropsTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomPropsTaskResultResponse), nil
	}
}

type CreateGenerationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGenerationTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGenerationTaskInvoker) Invoke() (*model.CreateGenerationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGenerationTaskResponse), nil
	}
}

type ShowGenerationTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGenerationTaskResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGenerationTaskResultInvoker) Invoke() (*model.ShowGenerationTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGenerationTaskResultResponse), nil
	}
}

type CheckDrugLigandDifferenceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckDrugLigandDifferenceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckDrugLigandDifferenceInvoker) Invoke() (*model.CheckDrugLigandDifferenceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckDrugLigandDifferenceResponse), nil
	}
}

type CreateDrugLigandInteraction2dSvgInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDrugLigandInteraction2dSvgInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDrugLigandInteraction2dSvgInvoker) Invoke() (*model.CreateDrugLigandInteraction2dSvgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDrugLigandInteraction2dSvgResponse), nil
	}
}

type CreateDrugLigandPreviewTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDrugLigandPreviewTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDrugLigandPreviewTaskInvoker) Invoke() (*model.CreateDrugLigandPreviewTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDrugLigandPreviewTaskResponse), nil
	}
}

type CreateDrugLigandSdfInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDrugLigandSdfInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDrugLigandSdfInvoker) Invoke() (*model.CreateDrugLigandSdfResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDrugLigandSdfResponse), nil
	}
}

type CreateDrugLigandSimilarityGraphTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDrugLigandSimilarityGraphTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDrugLigandSimilarityGraphTaskInvoker) Invoke() (*model.CreateDrugLigandSimilarityGraphTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDrugLigandSimilarityGraphTaskResponse), nil
	}
}

type CreateDrugLigandSvgInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDrugLigandSvgInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDrugLigandSvgInvoker) Invoke() (*model.CreateDrugLigandSvgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDrugLigandSvgResponse), nil
	}
}

type DeleteDrugLigandPreviewTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDrugLigandPreviewTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDrugLigandPreviewTaskInvoker) Invoke() (*model.DeleteDrugLigandPreviewTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDrugLigandPreviewTaskResponse), nil
	}
}

type DeleteDrugLigandSimilarityGraphTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDrugLigandSimilarityGraphTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDrugLigandSimilarityGraphTaskInvoker) Invoke() (*model.DeleteDrugLigandSimilarityGraphTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDrugLigandSimilarityGraphTaskResponse), nil
	}
}

type ParseDrugReceptorInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ParseDrugReceptorInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ParseDrugReceptorInfoInvoker) Invoke() (*model.ParseDrugReceptorInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParseDrugReceptorInfoResponse), nil
	}
}

type RecognizeDrugReceptorPocketInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeDrugReceptorPocketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeDrugReceptorPocketInvoker) Invoke() (*model.RecognizeDrugReceptorPocketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeDrugReceptorPocketResponse), nil
	}
}

type RunDrugLigandToSmilesConversionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDrugLigandToSmilesConversionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunDrugLigandToSmilesConversionInvoker) Invoke() (*model.RunDrugLigandToSmilesConversionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDrugLigandToSmilesConversionResponse), nil
	}
}

type RunDrugReceptorPreprocessInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDrugReceptorPreprocessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunDrugReceptorPreprocessInvoker) Invoke() (*model.RunDrugReceptorPreprocessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDrugReceptorPreprocessResponse), nil
	}
}

type ShowDrugLigandPreviewTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDrugLigandPreviewTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDrugLigandPreviewTaskInvoker) Invoke() (*model.ShowDrugLigandPreviewTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDrugLigandPreviewTaskResponse), nil
	}
}

type ShowDrugLigandSimilarityGraphTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDrugLigandSimilarityGraphTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDrugLigandSimilarityGraphTaskInvoker) Invoke() (*model.ShowDrugLigandSimilarityGraphTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDrugLigandSimilarityGraphTaskResponse), nil
	}
}

type DownloadDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadDataInvoker) Invoke() (*model.DownloadDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadDataResponse), nil
	}
}

type ShowOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOverviewInvoker) Invoke() (*model.ShowOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOverviewResponse), nil
	}
}

type CleanNextflowCacheInvoker struct {
	*invoker.BaseInvoker
}

func (i *CleanNextflowCacheInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CleanNextflowCacheInvoker) Invoke() (*model.CleanNextflowCacheResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CleanNextflowCacheResponse), nil
	}
}

type CreateNextflowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNextflowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNextflowJobInvoker) Invoke() (*model.CreateNextflowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNextflowJobResponse), nil
	}
}

type CreateNextflowWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNextflowWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNextflowWorkflowInvoker) Invoke() (*model.CreateNextflowWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNextflowWorkflowResponse), nil
	}
}

type DeleteNextflowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNextflowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNextflowJobInvoker) Invoke() (*model.DeleteNextflowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNextflowJobResponse), nil
	}
}

type DeleteNextflowWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNextflowWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNextflowWorkflowInvoker) Invoke() (*model.DeleteNextflowWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNextflowWorkflowResponse), nil
	}
}

type InstallNextflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *InstallNextflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InstallNextflowInvoker) Invoke() (*model.InstallNextflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstallNextflowResponse), nil
	}
}

type ListNextflowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNextflowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNextflowJobInvoker) Invoke() (*model.ListNextflowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNextflowJobResponse), nil
	}
}

type ListNextflowTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNextflowTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNextflowTaskInvoker) Invoke() (*model.ListNextflowTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNextflowTaskResponse), nil
	}
}

type ListNextflowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNextflowVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNextflowVersionInvoker) Invoke() (*model.ListNextflowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNextflowVersionResponse), nil
	}
}

type ListNextflowWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNextflowWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNextflowWorkflowInvoker) Invoke() (*model.ListNextflowWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNextflowWorkflowResponse), nil
	}
}

type RetryNextflowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryNextflowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryNextflowJobInvoker) Invoke() (*model.RetryNextflowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryNextflowJobResponse), nil
	}
}

type ShowNextflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNextflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNextflowInvoker) Invoke() (*model.ShowNextflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNextflowResponse), nil
	}
}

type ShowNextflowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNextflowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNextflowJobInvoker) Invoke() (*model.ShowNextflowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNextflowJobResponse), nil
	}
}

type ShowNextflowJobLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNextflowJobLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNextflowJobLogInvoker) Invoke() (*model.ShowNextflowJobLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNextflowJobLogResponse), nil
	}
}

type ShowNextflowJobReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNextflowJobReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNextflowJobReportsInvoker) Invoke() (*model.ShowNextflowJobReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNextflowJobReportsResponse), nil
	}
}

type ShowNextflowTaskDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNextflowTaskDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNextflowTaskDetailInvoker) Invoke() (*model.ShowNextflowTaskDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNextflowTaskDetailResponse), nil
	}
}

type ShowNextflowTaskLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNextflowTaskLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNextflowTaskLogInvoker) Invoke() (*model.ShowNextflowTaskLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNextflowTaskLogResponse), nil
	}
}

type ShowNextflowWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNextflowWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNextflowWorkflowInvoker) Invoke() (*model.ShowNextflowWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNextflowWorkflowResponse), nil
	}
}

type StopNextflowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopNextflowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopNextflowJobInvoker) Invoke() (*model.StopNextflowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopNextflowJobResponse), nil
	}
}

type UninstallNextflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UninstallNextflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UninstallNextflowInvoker) Invoke() (*model.UninstallNextflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UninstallNextflowResponse), nil
	}
}

type UpdateNextflowWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNextflowWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNextflowWorkflowInvoker) Invoke() (*model.UpdateNextflowWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNextflowWorkflowResponse), nil
	}
}

type CreateNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNotebookInvoker) Invoke() (*model.CreateNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNotebookResponse), nil
	}
}

type DeleteNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNotebookInvoker) Invoke() (*model.DeleteNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNotebookResponse), nil
	}
}

type ListNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotebookInvoker) Invoke() (*model.ListNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotebookResponse), nil
	}
}

type ListNotebookToolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotebookToolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotebookToolInvoker) Invoke() (*model.ListNotebookToolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotebookToolResponse), nil
	}
}

type ShowNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNotebookInvoker) Invoke() (*model.ShowNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotebookResponse), nil
	}
}

type ShowNotebookTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotebookTokenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNotebookTokenInvoker) Invoke() (*model.ShowNotebookTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotebookTokenResponse), nil
	}
}

type StopOrStartNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopOrStartNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopOrStartNotebookInvoker) Invoke() (*model.StopOrStartNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopOrStartNotebookResponse), nil
	}
}

type UpdateNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNotebookInvoker) Invoke() (*model.UpdateNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNotebookResponse), nil
	}
}

type ListObsBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObsBucketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListObsBucketInvoker) Invoke() (*model.ListObsBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObsBucketResponse), nil
	}
}

type ListObsBucketObjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObsBucketObjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListObsBucketObjectInvoker) Invoke() (*model.ListObsBucketObjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObsBucketObjectResponse), nil
	}
}

type CreateOptimizationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOptimizationTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOptimizationTaskInvoker) Invoke() (*model.CreateOptimizationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOptimizationTaskResponse), nil
	}
}

type ShowOptimizationTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOptimizationTaskResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOptimizationTaskResultInvoker) Invoke() (*model.ShowOptimizationTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOptimizationTaskResultResponse), nil
	}
}

type CreateSearchTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSearchTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSearchTaskInvoker) Invoke() (*model.CreateSearchTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSearchTaskResponse), nil
	}
}

type ShowSearchTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSearchTaskResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSearchTaskResultInvoker) Invoke() (*model.ShowSearchTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSearchTaskResultResponse), nil
	}
}

type CreateSynthesisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSynthesisTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSynthesisTaskInvoker) Invoke() (*model.CreateSynthesisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSynthesisTaskResponse), nil
	}
}

type ShowSynthesisTaskResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSynthesisTaskResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSynthesisTaskResultInvoker) Invoke() (*model.ShowSynthesisTaskResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSynthesisTaskResultResponse), nil
	}
}

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

type CreateClusteringJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusteringJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusteringJobInvoker) Invoke() (*model.CreateClusteringJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusteringJobResponse), nil
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

type CreateComputingClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateComputingClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateComputingClusterInvoker) Invoke() (*model.CreateComputingClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateComputingClusterResponse), nil
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

type CreateDrugModelResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDrugModelResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDrugModelResourceInvoker) Invoke() (*model.CreateDrugModelResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDrugModelResourceResponse), nil
	}
}

type CreateFavoriteInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFavoriteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFavoriteInvoker) Invoke() (*model.CreateFavoriteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFavoriteResponse), nil
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

type CreateMolDockingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMolDockingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMolDockingJobInvoker) Invoke() (*model.CreateMolDockingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMolDockingJobResponse), nil
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

type DeleteComputingClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteComputingClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteComputingClusterInvoker) Invoke() (*model.DeleteComputingClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteComputingClusterResponse), nil
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

type DeleteDrugModelResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDrugModelResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDrugModelResourceInvoker) Invoke() (*model.DeleteDrugModelResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDrugModelResourceResponse), nil
	}
}

type DeleteFavoriteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFavoriteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFavoriteInvoker) Invoke() (*model.DeleteFavoriteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFavoriteResponse), nil
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

type InitializePlatformInvoker struct {
	*invoker.BaseInvoker
}

func (i *InitializePlatformInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InitializePlatformInvoker) Invoke() (*model.InitializePlatformResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InitializePlatformResponse), nil
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

type ListCceClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCceClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCceClusterInvoker) Invoke() (*model.ListCceClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCceClusterResponse), nil
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

type ListClusterInstallStepInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterInstallStepInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterInstallStepInvoker) Invoke() (*model.ListClusterInstallStepResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterInstallStepResponse), nil
	}
}

type ListComputingClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComputingClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListComputingClusterInvoker) Invoke() (*model.ListComputingClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComputingClusterResponse), nil
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

type ListDrugModelResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDrugModelResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDrugModelResourceInvoker) Invoke() (*model.ListDrugModelResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDrugModelResourceResponse), nil
	}
}

type ListFavoriteInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFavoriteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFavoriteInvoker) Invoke() (*model.ListFavoriteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFavoriteResponse), nil
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

type ListProjectStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectStatisticsInvoker) Invoke() (*model.ListProjectStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectStatisticsResponse), nil
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

type ListSfsTurbosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSfsTurbosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSfsTurbosInvoker) Invoke() (*model.ListSfsTurbosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSfsTurbosResponse), nil
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

type ListUserAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserAppInvoker) Invoke() (*model.ListUserAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserAppResponse), nil
	}
}

type ListUserDrugJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserDrugJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserDrugJobInvoker) Invoke() (*model.ListUserDrugJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserDrugJobResponse), nil
	}
}

type ListUserImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserImageInvoker) Invoke() (*model.ListUserImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserImageResponse), nil
	}
}

type ListUserJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserJobInvoker) Invoke() (*model.ListUserJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserJobResponse), nil
	}
}

type ListUserWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserWorkflowInvoker) Invoke() (*model.ListUserWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserWorkflowResponse), nil
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

type RunFormatConverterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunFormatConverterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunFormatConverterInvoker) Invoke() (*model.RunFormatConverterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunFormatConverterResponse), nil
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

type ShowAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgencyInvoker) Invoke() (*model.ShowAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgencyResponse), nil
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

type ShowClusteringJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusteringJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusteringJobInvoker) Invoke() (*model.ShowClusteringJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusteringJobResponse), nil
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

type UpdateAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAgencyInvoker) Invoke() (*model.UpdateAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAgencyResponse), nil
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

type UpdateTopProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTopProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTopProjectInvoker) Invoke() (*model.UpdateTopProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTopProjectResponse), nil
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

type ListUserNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserNotebookInvoker) Invoke() (*model.ListUserNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserNotebookResponse), nil
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

type DownloadPublicDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadPublicDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadPublicDataInvoker) Invoke() (*model.DownloadPublicDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadPublicDataResponse), nil
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

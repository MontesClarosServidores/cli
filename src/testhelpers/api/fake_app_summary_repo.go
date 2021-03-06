package api

import (
	"cf/errors"
	"cf/models"
)

type FakeAppSummaryRepo struct {
	GetSummariesInCurrentSpaceApps []models.Application

	GetSummaryErrorCode string
	GetSummaryAppGuid   string
	GetSummarySummary   models.Application
}

func (repo *FakeAppSummaryRepo) GetSummariesInCurrentSpace() (apps []models.Application, apiErr error) {
	apps = repo.GetSummariesInCurrentSpaceApps
	return
}

func (repo *FakeAppSummaryRepo) GetSummary(appGuid string) (summary models.Application, apiErr error) {
	repo.GetSummaryAppGuid = appGuid
	summary = repo.GetSummarySummary

	if repo.GetSummaryErrorCode != "" {
		apiErr = errors.NewHttpError(400, repo.GetSummaryErrorCode, "Error")
	}

	return
}

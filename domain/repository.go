package domain

import (
	"time"
)

type Repository interface {
	Ping() error
	Close() error

	ListCollections() ([]*CollectionItem, error)
	CreateCollection(c *Collection) error
	RetrieveCollection(id string) (*Collection, error)
	UpdateCollection(c *Collection) error
	DeleteCollection(id string) error

	ListVariables(collectionID string) ([]*VariableItem, error)
	MapVariables(collectionID string) (map[string]string, error)
	CreateVariable(v *Variable) error
	RetrieveVariable(id string) (*Variable, error)
	UpdateVariable(v *Variable) error
	DeleteVariable(id string) error

	ListJobs(collectionID string, fields []string) ([]*JobItem, error)
	CreateJob(j *JobDefinition) error
	RetrieveJob(id string) (*JobDefinition, error)
	UpdateJob(j *JobDefinition) error
	DeleteJob(id string) error

	RetrieveJobStatus(id string) (*JobStatus, error)
	ListLeftOverJobs() ([]string, error)
	ResetJobStatus(id string) error

	ListJobHistory(id string) ([]*JobHistory, error)
	DeleteJobHistory(id string, before time.Time) error

	AcquireJob(id string, deadline time.Duration) error
	AddJobHistory(*JobHistory) error
}

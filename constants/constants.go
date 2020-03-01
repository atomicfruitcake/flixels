package constants

// AppPort HTTP port where the application is run on
const AppPort = "8080"

// Job data used to identify and transfer jobs handles by flixels
type Job struct {
	AssetID string
	JobID   string
	Service string
	Status  int
}

// JobReq data request format used when querying job statuses
type JobReq struct {
	JobID string
}

const (
	Pending    int = 0
	Processing int = 1
	Success    int = 2
	Error      int = 3
)

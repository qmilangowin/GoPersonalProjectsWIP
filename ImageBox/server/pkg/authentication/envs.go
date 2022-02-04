package authentication

import (
	"os"
)

type GCPEnvs struct {
	ProjectID string
	Bucket    string
}

func GetEnvs() *GCPEnvs {

	pid := os.Getenv("GCLOUD_PROJECT")
	bucket := os.Getenv("BUCKET")

	gcp := GCPEnvs{
		ProjectID: pid,
		Bucket:    bucket,
	}

	return &gcp
}

//export GCLOUD_PROJECT=<>
//export BUCKET=<>

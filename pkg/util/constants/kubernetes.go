package constants

import "time"

// General kubernetes-related constants
const (
	K8sAPIPath                  = "/apis"
	StatefulsetDeletionInterval = 1 * time.Second
	StatefulsetDeletionTimeout  = 30 * time.Second

	QueueResyncPeriodPod  = 5 * time.Minute
	QueueResyncPeriodTPR  = 5 * time.Minute
	QueueResyncPeriodNode = 5 * time.Minute
)

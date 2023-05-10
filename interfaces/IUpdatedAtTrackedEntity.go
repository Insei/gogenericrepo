package interfaces

import "time"

type IUpdatedAtTrackedEntity interface {
	//GetUpdatedAt time
	//
	//Return
	//	*time.Time - update time
	GetUpdatedAt() *time.Time
}

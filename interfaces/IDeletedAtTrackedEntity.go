package interfaces

import "time"

type IDeletedAtTrackedEntity interface {
	//GetDeletedAt time
	//
	//Return
	//	time.Time - update time
	GetDeletedAt() *time.Time
}

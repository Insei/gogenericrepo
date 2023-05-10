package interfaces

import (
	"time"
)

// IEntity default interface that must be implemented by each entity
type IEntity[IDType comparable] interface {
	//GetID of the entity
	//
	//Return
	//	IDType - entity id
	GetID() IDType
	//GetCreatedAt time
	//
	//Return
	//	uuid.UUID - create time
	GetCreatedAt() time.Time
}

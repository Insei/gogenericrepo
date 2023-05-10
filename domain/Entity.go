package domain

import (
	"github.com/google/uuid"
	"time"
)

//BaseEntity structure
type BaseEntity[IDType comparable] struct {
	//	ID - entity identifier
	ID IDType `gorm:"primary_key;"`
	//	CreatedAt - time when the entity was created
	CreatedAt time.Time `gorm:"index"`
}

//GetID of the entity
func (e BaseEntity[IDType]) GetID() IDType {
	return e.ID
}

//GetCreatedAt time
func (e BaseEntity[IDType]) GetCreatedAt() time.Time {
	return e.CreatedAt
}

type UpdatedAtTrackedEntity[IDType comparable] struct {
	BaseEntity[IDType]
	//	UpdatedAt - time when the entity was updated
	UpdatedAt *time.Time `gorm:"index"`
}

//GetUpdatedAt time
func (e UpdatedAtTrackedEntity[IDType]) GetUpdatedAt() *time.Time {
	return e.UpdatedAt
}

type DeletedAtTrackedEntity[IDType comparable] struct {
	BaseEntity[IDType]
	//	DeletedAt - time when the entity was deleted
	DeletedAt *time.Time `gorm:"index"`
}

//GetDeletedAt time
func (e DeletedAtTrackedEntity[IDType]) GetDeletedAt() *time.Time {
	return e.DeletedAt
}

type DateTrackedEntity[IDType comparable] struct {
	BaseEntity[IDType]
	//	UpdatedAt - time when the entity was updated
	UpdatedAt *time.Time `gorm:"index"`
	//	DeletedAt - time when the entity was deleted
	DeletedAt *time.Time `gorm:"index"`
}

//GetUpdatedAt time
func (e DateTrackedEntity[IDType]) GetUpdatedAt() *time.Time {
	return e.UpdatedAt
}

//GetDeletedAt time
func (e DateTrackedEntity[IDType]) GetDeletedAt() *time.Time {
	return e.DeletedAt
}

//DateTrackedEntityUUID date tracked entity structure where ID type is uuid.UUID
type DateTrackedEntityUUID struct {
	DateTrackedEntity[uuid.UUID]
	//	ID - entity identifier
	ID uuid.UUID `gorm:"primary_key; type:varchar(36)"`
}

//GetID of the entity
func (e DateTrackedEntityUUID) GetID() uuid.UUID {
	return e.ID
}

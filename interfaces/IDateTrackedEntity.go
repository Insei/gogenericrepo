package interfaces

type IDateTrackedEntity[IDType comparable] interface {
	IEntity[IDType]
	IUpdatedAtTrackedEntity
	IDeletedAtTrackedEntity
}

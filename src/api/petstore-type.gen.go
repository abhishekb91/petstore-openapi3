// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

// Error defines model for Error.
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// Pet defines model for Pet.
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/PetRequest)
	PetRequest `yaml:",inline"`
	// Embedded fields due to inline allOf schema

	// Unique id of the pet
	Id int64 `json:"id"`
}

// PetRequest defines model for PetRequest.
type PetRequest struct {
	Name string `json:"name"`

	// pet status in the store
	Status *string `json:"status,omitempty"`
}

// AddPetJSONBody defines parameters for AddPet.
type AddPetJSONBody PetRequest

// UpdatePetByIdJSONBody defines parameters for UpdatePetById.
type UpdatePetByIdJSONBody PetRequest

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody AddPetJSONBody

// UpdatePetByIdJSONRequestBody defines body for UpdatePetById for application/json ContentType.
type UpdatePetByIdJSONRequestBody UpdatePetByIdJSONBody

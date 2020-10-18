package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/satori/go.uuid"
)

type ModelInfo struct {
	ResourceType     reflect.Type
	IDFieldName      string
	IDFieldType      reflect.Type
	IDFieldZeroValue reflect.Value
}

func GetResourceModelInfo(m Model) ModelInfo {
	return ModelInfo{
		ResourceType:     reflect.TypeOf(m),
		IDFieldName:      m.GetIDFieldName(),
		IDFieldType:      reflect.TypeOf(m.GetID()),
		IDFieldZeroValue: reflect.Zero(reflect.TypeOf(m.GetID())),
	}
}

type Model interface {
	GetID() uint
	GetIDFieldName() string
	GetPublicIDFieldName() string
	GetOwnerIDFieldName() string
}

type BaseModel struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (b BaseModel) GetID() uint {
	return b.ID
}

func (b BaseModel) GetIDFieldName() string {
	return "ID"
}

func (b BaseModel) GetPublicIDFieldName() string {
	return ""
}

func (b BaseModel) GetOwnerIDFieldName() string {
	return ""
}

type BaseModelWithPublicID struct {
	BaseModel
	PublicID uuid.UUID `json:"id"`
}

func (bid BaseModelWithPublicID) GetPublicIDFieldName() string {
	return "PublicID"
}

type ValidationErrors []string

func (v ValidationErrors) Error() string {
	verrs, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Errorf("marshalling of validation errors failed: %v", err))
	}

	return string(verrs)
}

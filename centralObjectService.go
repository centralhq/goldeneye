package main

import (
	"errors"
	"net/http"
  "encoding/json"
  "log"
  "bytes"
  "fmt"
)

type CentralObjectService struct {
  updateBuilder CentralObjectBuilder
}

type ObjectType int

const (
  UPDATE ObjectType = iota
  DELETE
  CREATE
)

func NewCentralObjectService (updateBuilder CentralObjectBuilder) *CentralObjectService {
  return &CentralObjectService{
	updateBuilder: updateBuilder,
  }
}

func (c *CentralObjectService) create(payload interface{}, objectId string, opType string, uuid string) error {
  return c.applyOperation(opType, objectId, uuid, payload, false)
}

func (c *CentralObjectService) update(payload interface{}, objectId string, opType string, uuid string) error {
  return c.applyOperation(opType, objectId, uuid, payload, false)
}

func (c *CentralObjectService) delete(payload interface{}, objectId string, opType string, uuid string) error {
  return c.applyOperation(opType, objectId, uuid, payload, true)
}

func (c *CentralObjectService) applyOperation(opType string, objectId string, uuid string, payload interface{}, isDeleted bool) error {
	conflictId, err := c.createConflictId(opType, objectId)
	if err != nil {
		return err
	}

	object := c.buildCentralObject(opType, uuid, conflictId, isDeleted, payload)

	c.resolveObject(object)

	return nil
}

func (c *CentralObjectService) buildCentralObject(opType string, uuid string, conflictId string, isDeleted bool, payload interface{}) *CentralObject {
  return c.updateBuilder.withOpType(opType).
  withUuid(uuid).
  withConflictId(conflictId).
  withIsDeleted(isDeleted).
  withPayload(payload).
  build()
}

func (c *CentralObjectService) createConflictId(opType string, objectId string) (string, error) {
  if len(opType) > 0 && len(objectId) > 0 {
    return opType + "_" + objectId, nil
  }
  return "", errors.New("Invalid input")
}

/**
* resolveObject sends off transformed object to Central Service.
*/
func (c *CentralObjectService) resolveObject(object *CentralObject) error {
  
  objectByte, err := json.Marshal(object)
  if err != nil {
    log.Println(err)
  }
  buffer := bytes.NewReader(objectByte)
  request, err := http.NewRequest(http.MethodPost, "http://example.com/send", buffer)
  
  if err != nil {
    fmt.Printf("client: could not create request: %s\n", err)
    return err
  }

  response, err := http.DefaultClient.Do(request)

  if err != nil {
    fmt.Printf("client: could not read response body: %s\n", err)
    return err
  }

  fmt.Printf("client: response body: %s\n", response.Body)
  return nil
}
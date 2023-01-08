package main

type CentralObjectBuilder struct {
	opType 	    string
	uuid 		string
	conflictId	string
	isDeleted   bool
	payload     interface{}
  }
  
  func NewCentralObjectBuilder() *CentralObjectBuilder {
	return &CentralObjectBuilder{}
  }
  
  func (c *CentralObjectBuilder) withOpType(opType string) *CentralObjectBuilder {
	c.opType = opType
	return c
  }
  
  func (c *CentralObjectBuilder) withUuid(uuid string) *CentralObjectBuilder {
	c.uuid = uuid
	return c
  }
  
  func (c *CentralObjectBuilder) withConflictId(conflictId string) *CentralObjectBuilder {
	c.conflictId = conflictId
	return c
  }
  
  func (c *CentralObjectBuilder) withIsDeleted(isDeleted bool) *CentralObjectBuilder {
	c.isDeleted = isDeleted
	return c
  }
  
  func (c *CentralObjectBuilder) withPayload(payload interface{}) *CentralObjectBuilder {
	c.payload = payload
	return c
  }
  
  func (c *CentralObjectBuilder) build() *CentralObject {
	// TODO: add validation validate

    return &CentralObject{
	  OpType: c.opType,
	  Uuid: c.uuid,
	  ConflictId: c.conflictId,
	  IsDeleted: c.isDeleted,
	  Payload: c.payload,	
	}
  }
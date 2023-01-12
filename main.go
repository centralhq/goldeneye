package main

type CentralObject struct {
  OpType 		string		`json:"opType"`
  Uuid		    string		`json:"uuId"`
  ConflictId 	string  	`json:"conflictId"`
  IsDeleted 	bool		`json:"isDeleted"`
  Payload 	    interface{}	`json:"payload"`
}
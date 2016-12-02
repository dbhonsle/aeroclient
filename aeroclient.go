package main

import (
	//"crypto/rsa"
	//"flag"
	"fmt"
	as "github.com/aerospike/aerospike-client-go"
	"github.com/pborman/uuid"
	//"io/ioutil"
	//"os"
	//"strings"
	//"log"
	//"encoding/json"
	//"errors"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/googollee/go-socket.io"
	//"github.com/julienschmidt/httprouter"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//"math/rand"
	//"net/http"
	//"strconv"
	"time"
)

type ContactStruct struct {
	Name string `as:"name"` // alias the field to a
	//Self *ContactStruct    `as:"-"`  // will not persist the field
	CreateDate time.Time `as:"createDate"`
}

func main() {
	client, err := as.NewClient("192.168.2.70", 3000)

	if err == nil {
		fmt.Println("Connected To Aerospike Server Successfully")
	} else {
		fmt.Println(err)
	}

	// NameSpace, Set, Key-> is uuid
	key, err := as.NewKey("test", "intellidoccontactsset",
		[]byte(uuid.NewRandom()))

	/*
		bin1 := as.NewBin("name", "Deepak Bhonsle")
		bin2 := as.NewBin("email", "dsbhonsle@gmail.com")

		// Write a record
		err = client.PutBins(nil, key, bin1, bin2)

		// Read a record
		record, err := client.Get(nil, key)

		fmt.Println(record.Bins["name"].(string))
	*/
	contact := &ContactStruct{Name: "Devayani Bhonsle", CreateDate: time.Now()}

	err = client.PutObject(nil, key, contact)

	rcontact := &ContactStruct{}

	err = client.GetObject(nil, key, rcontact)

	fmt.Println(rcontact.Name)
	fmt.Println(rcontact.CreateDate)

	client.Delete(nil, key)

	client.Close()
}

/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	// "strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("example_cc0")

type User struct {
    Id          string
    Name        string
    Lastname    string
    Email       string
    Amount      float64
		LoanId			string
}

type bank struct {
    Id          string
    Name        string
    Year        int
    Country     string
    Countries   []string
    Users       []User
}

type Transaction struct {
    Id          string
    Date        string
    Id_sender   string
    Id_receiver string
    Amount      float64
}

type Loan   struct {
    Id          string
    Date_start  string
    Date_end    string
    Rate        float64
    Interest    float64
    Num_rates   int
    Num_paId    int
    Amount      float64
}

// type tutor struct {
// 	Id      string
// 	Name    string
// 	Surname string
// }
//
// type tutorial struct {
// 	Id     string
// 	Name   string
// 	Tutors []string
// }
//
// // Global variables for Id
// var tutorId int
// var tutorialId int

var userIdCount int
var transactionIdCount int
var loanIdCount int

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Init ###########")

    // 2 banke, 3 korisnika, 4 kredita, 4 transakcije

    var loan1 = Loan{"l1", "010101", "010701", 500, 1.5, 6, 3, 3000}
    var loan2 = Loan{"l2", "010105", "010705", 1000, 1.5, 6, 2, 6000}
    var loan3 = Loan{"l3", "150302", "150702", 500, 1.5, 4, 2, 2000}
    var loan4 = Loan{"l4", "200105", "201205", 500, 1.5, 11, 3, 5500}


		loanIdCount = 5

    var user1 = User{"u1", "Nikola", "Malencic", "nmalencic@gmail.com", 1000, "l1"}
    var user2 = User{"u2", "Marko", "Markovic", "nmalencic@gmail.com", 20000, "l2"}
    var user3 = User{"u3", "Cvetko", "Cvetkovic", "nmalencic@gmail.com", 5000, "l3"}
    var user4 = User{"u4", "Jovan", "Jovanovic", "nmalencic@gmail.com", 10000, "l4"}

		userIdCount = 5

    var bank1countries = make([]string, 0, 10)
    bank1countries = append(bank1countries, "Serbia")
    bank1countries = append(bank1countries, "USA")

    var bank2countries = make([]string, 0, 10)
    bank2countries = append(bank2countries, "Serbia")
    bank2countries = append(bank2countries, "USA")

    var bank1users = make([]User, 0, 10)
    bank1users = append(bank1users, user1)
    bank1users = append(bank1users, user2)

    var bank2users = make([]User, 0, 10)
    bank2users = append(bank2users, user3)
    bank2users = append(bank2users, user4)

    var bank1 = bank{"b1", "bank1", 1996, "Serbia", bank1countries, bank1users}
    var bank2 = bank{"b2", "bank2", 2000, "Serbia", bank2countries, bank2users}

    var transaction1 = Transaction{"t1", "050505", "u1", "u2", 50}
    var transaction2 = Transaction{"t2", "061010", "u1", "u3", 50}
    var transaction3 = Transaction{"t3", "050509", "u3", "u2", 50}
    var transaction4 = Transaction{"t4", "221015", "u4", "u1", 1000}

		transactionIdCount = 5

    ajson, _ := json.Marshal(user1)
    err := stub.PutState(user1.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(user2)
    err = stub.PutState(user2.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(user3)
    err = stub.PutState(user3.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(user4)
    err = stub.PutState(user4.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(transaction1)
    err = stub.PutState(transaction1.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(transaction2)
    err = stub.PutState(transaction2.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(transaction3)
    err = stub.PutState(transaction3.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(transaction4)
    err = stub.PutState(transaction4.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(loan1)
    err = stub.PutState(loan1.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(loan2)
    err = stub.PutState(loan2.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(loan3)
    err = stub.PutState(loan3.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(loan4)
    err = stub.PutState(loan4.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(bank1)
		// test := string(ajson[:])

		// return shim.Error("bank1.Id je " + bank1.Id + " a ajson tu je " + test)
    err = stub.PutState(bank1.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }

    ajson, _ = json.Marshal(bank2)
    err = stub.PutState(bank2.Id, ajson)
    if err != nil {
        return shim.Error(err.Error())
    }


	/*var tutor1 = tutor{"tu1", "John", "Doe"}
	var tutor2 = tutor{"tu2", "Michel", "Green"}
	var tutor3 = tutor{"tu3", "Jova", "Jovanovic"}
	tutorId = 4

	var tutorsFor1 = make([]string, 0, 20)
	tutorsFor1 = append(tutorsFor1, tutor1.Id)
	tutorsFor1 = append(tutorsFor1, tutor2.Id)
	 var tutorial1 = tutorial{"t1","Blockcahin tutorial",tutorsFor1}
	var tutorial1 = tutorial{"t1", "Blockcahin tutorial", tutorsFor1}

	var tutorsFor2 = make([]string, 0, 20)
	tutorsFor2 = append(tutorsFor2, tutor3.Id)
	var tutorial2 = tutorial{"t2", "Spark tutorial", tutorsFor2}
	tutorialId = 3

	// Write the state to the ledger
	ajson, _ := json.Marshal(tutor1)
	err := stub.PutState("tu1", ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(tutor2)
	err = stub.PutState("tu2", ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	ajson, _ = json.Marshal(tutor3)
	err = stub.PutState("tu3", ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(tutorial1)
	err = stub.PutState("t1", ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(tutorial2)
	err = stub.PutState("t2", ajson)
	if err != nil {
		return shim.Error(err.Error())
	}*/

	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "delete" {
		return t.delete(stub, args)
	}
	if function == "query" {
		return t.query(stub, args)
	}
	if function == "addUser" {
		return t.addUser(stub, args)
	}
	if function == "getLoan"{
		return t.getLoan(stub, args)
	}
	if function == "payLoan" {
		return t.payLoan(stub, args)
	}

	if function == "addTransaction" {
		return t.addTransaction(stub, args)
	}
	// if function == "addTutorial" {
	// 	return t.addTutorial(stub, args)
	// }
	// if function == "addTutor" {
	// 	return t.addTutor(stub, args)
	// }
	// if function == "addTutorToTutorial" {
	// 	return t.addTutorToTutorial(stub, args)
	// }
	// if function == "removeTutorFromTutorial" {
	// 	return t.removeTutorFromTutorial(stub, args)
	// }

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}
// usage: addUser(name, lastname, email, amount, bankId)
func (t *SimpleChaincode) addUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments, expected 5, got " + strconv.Itoa(len(args)))
	}

	name := args[0]
	lastname := args[1]
	email := args[2]
	amount, err := strconv.ParseFloat(args[3], 64)
	bankId := args[4]


	Id := "u"+ strconv.Itoa(userIdCount)


	userIdCount = userIdCount + 1

	var user = User{Id, name, lastname, email, amount, "0"}

	logger.Debug(user)

	ajson, _ := json.Marshal(user)

	err = stub.PutState(user.Id, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	bankBytes, err := stub.GetState(bankId)
	if err != nil {
		return shim.Error("Failed to get bank state")
	}

	if bankBytes == nil || len(bankBytes) == 0 {
		return shim.Error("Bank doesn't exist!")
	}

	bank := bank{}
	err = json.Unmarshal(bankBytes, &bank)
	if err != nil {
		return shim.Error(err.Error())
	}

	bank.Users = append(bank.Users, user)

	// test := string(bankBytes[:])
	// len := strconv.Itoa(len(bankBytes))
	// 	return shim.Error("bankBytes je " + test + " a len(bankBytes) je " + len)
	bankBytes, _ = json.Marshal(bank)

	err = stub.PutState(bank.Id, bankBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}
// usage addTransaction(senderId, receiverId, amount, minus)
func (t *SimpleChaincode) addTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Usage is args: sender, receiver, amount, minus")
	}

	senderId := args[0]
	receiverId := args[1]
	amount, _ := strconv.ParseFloat(args[2], 64)
	minus, _ := strconv.Atoi(args[3])

	if minus != 0 && minus != 1 {
		return shim.Error("Minus argument has to be 0 or 1")
	}

	sender := User{}
	receiver := User{}

	senderBytes, err := stub.GetState(senderId)
	if err != nil {
		return shim.Error(err.Error() + "1")
	}

	err = json.Unmarshal(senderBytes, &sender)
	if err != nil {
		return shim.Error(err.Error() + "2")
	}

	receiverBytes, err := stub.GetState(receiverId)
	if err != nil {
		return shim.Error(err.Error() + "3")
	}

	err = json.Unmarshal(receiverBytes, &receiver)
	if err != nil {
		return shim.Error(err.Error() + "4")
	}

if sender.Amount < amount {
	if minus == 1 {
		// minus logic
		if sender.Amount < 0 {
			return shim.Error("Already in minus!")
		} else {
			sum := 0.0
			transNum := 0.0
			for i:=0;i<transactionIdCount-1;i++ {
				transactionBytes, err := stub.GetState("t" + strconv.Itoa(i+1))
				if err != nil {
					return shim.Error(err.Error() + "5")
					}

					transaction := Transaction{}

					err = json.Unmarshal(transactionBytes, &transaction)
					if err != nil {
						return shim.Error(err.Error() + "6")
						}

					if transaction.Id_receiver == sender.Id {
						sum = sum + transaction.Amount
						transNum = transNum + 1
					}
			}

			allowed := sum / transNum

			if sender.Amount - amount < (-allowed) {
				return shim.Error("Not enough money for minus")
			} else {
				transactionId := "t" + strconv.Itoa(transactionIdCount)
				transactionIdCount = transactionIdCount + 1
				transaction := Transaction{transactionId, time.Now().Format("2006-01-02 15:04:05"), sender.Id, receiver.Id, amount}
				sender.Amount = sender.Amount - amount
				receiver.Amount = receiver.Amount + amount

				ajson, _ := json.Marshal(sender)
				err = stub.PutState(sender.Id, ajson)
				if err != nil {
					return shim.Error(err.Error() + "7")
					}

				ajson, _ = json.Marshal(receiver)
				err = stub.PutState(receiver.Id, ajson)
				if err != nil {
					return shim.Error(err.Error() + "8")
					}

				ajson, _ = json.Marshal(transaction)
				err = stub.PutState(transaction.Id, ajson)
				if err != nil {
					return shim.Error(err.Error() + "9")
					}
			}
		}

	} else {
		return shim.Error("Not enough money")
	}
} else {
	transactionId := "t" + strconv.Itoa(transactionIdCount)
	transactionIdCount = transactionIdCount + 1
	transaction := Transaction{transactionId, time.Now().Format("2006-01-02 15:04:05"), sender.Id, receiver.Id, amount}
	sender.Amount = sender.Amount - amount
	receiver.Amount = receiver.Amount + amount

	ajson, _ := json.Marshal(sender)
	err = stub.PutState(sender.Id, ajson)
	if err != nil {
		return shim.Error(err.Error() + "10")
		}

	ajson, _ = json.Marshal(receiver)
	err = stub.PutState(receiver.Id, ajson)
	if err != nil {
		return shim.Error(err.Error() + "11")
		}

	ajson, _ = json.Marshal(transaction)
	err = stub.PutState(transaction.Id, ajson)
	if err != nil {
		return shim.Error(err.Error() + "12")
		}
}

return shim.Success(nil)




}

func (t *SimpleChaincode) getLoan(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Usage is args: userId, amount")
	}

	userId := args[0]
	amount, _ := strconv.ParseFloat(args[1], 64)

	user := User{}

	userBytes, err := stub.GetState(userId)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return shim.Error(err.Error())
	}

	if user.LoanId != "0" {
		return shim.Error("User already has a loan")
	}

	sum := 0.0
	transNum := 0.0
	for i:=0;i<transactionIdCount-1;i++ {
		transactionBytes, err := stub.GetState("t" + strconv.Itoa(i+1))
		if err != nil {
			return shim.Error(err.Error())
			}

			transaction := Transaction{}

			err = json.Unmarshal(transactionBytes, &transaction)
			if err != nil {
				return shim.Error(err.Error())
				}

			if transaction.Id_receiver == user.Id {
				sum = sum + transaction.Amount
				transNum = transNum + 1
			}
	}

	allowed := (sum/transNum)*5

	if amount > allowed {
		return shim.Error("You want too much money")
	}

	loanId := "l" + strconv.Itoa(loanIdCount)
	loanIdCount = loanIdCount + 1

	loan := Loan{loanId, time.Now().Format("2006-01-02 15:04:05"), time.Now().AddDate(0, 12, 0).Format("2006-01-02 15:04:05"), amount/12, 1.5, 12, 0, amount }

	ajson, _ := json.Marshal(loan)
	err = stub.PutState(loan.Id, ajson)
	if err != nil {
			return shim.Error(err.Error())
	}

	user.LoanId = loan.Id
	user.Amount = user.Amount + amount

	ajson, _ = json.Marshal(user)
	err = stub.PutState(user.Id, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *SimpleChaincode) payLoan(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Usage args: userId")
	}

	userId := args[0]

	user := User{}

	userBytes, err := stub.GetState(userId)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return shim.Error(err.Error())
	}

	loanId := user.LoanId

	loanBytes, err := stub.GetState(loanId)

	loan := Loan{}

	err = json.Unmarshal(loanBytes, &loan)

	if user.Amount < loan.Rate {
		return shim.Error("Not enough money")
	}

	user.Amount = user.Amount - loan.Rate * loan.Interest
	loan.Num_paId = loan.Num_paId + 1

	if loan.Num_paId == loan.Num_rates {
		user.LoanId = "0"
	}

	userBytes, _ = json.Marshal(user)
	err = stub.PutState(user.Id, userBytes)

	loanBytes, _ = json.Marshal(loan)
	err = stub.PutState(loan.Id, loanBytes)

	return shim.Success(nil)
}

// func (t *SimpleChaincode) addTutor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	var name, surname string // Entities
//
// 	if len(args) < 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
// 	}
//
// 	name = args[0]
// 	surname = args[1]
//
// 	tutorKey := "tu" + strconv.Itoa(tutorId)
// 	tutorId = tutorId + 1
// 	var newTutor = tutor{tutorKey, name, surname}
//
// 	ajson, _ := json.Marshal(newTutor)
// 	err := stub.PutState(newTutor.Id, ajson)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
//
// 	return shim.Success(nil)
// }
//
// func (t *SimpleChaincode) addTutorial(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	// TODO implement function
// 	// arg 0 - name, arg1,arg2,arg3,arg4... - tutorId (which is the same as tutorKey)
// 	// Check number of arguments
// 	// Check if tutors exist in ledger before adding them to tutorial
// 	var name, tutorKey string // Entities
//
// 	if len(args) < 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
// 	}
//
// 	name = args[0]
// 	var tutors = make([]string, 0, 20)
// 	for i := 1; i < len(args); i++ {
// 		tutorKey = args[i]
// 		tutorI, err := stub.GetState(tutorKey)
// 		logger.Info("Tutor " + tutorKey + "postoji")
// 		logger.Info(tutorI)
//
// 		if err != nil {
// 			jsonResp := "{\"Error\":\"Failed to get state for " + tutorKey + "\"}"
// 			return shim.Error(jsonResp)
// 		}
// 		if tutorI == nil || len(tutorI) == 0 {
// 			jsonResp := "{\"Error\":\" " + tutorKey + " does not exit " + "\"}"
// 			return shim.Error(jsonResp)
// 		}
//
// 		tutors = append(tutors, tutorKey)
// 	}
//
// 	tutorialKey := "t" + strconv.Itoa(tutorialId)
// 	tutorialId = tutorialId + 1
// 	var newTutorial = tutorial{tutorialKey, name, tutors}
//
// 	ajson, _ := json.Marshal(newTutorial)
// 	err := stub.PutState(newTutorial.Id, ajson)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
//
// 	return shim.Success(nil)
// }
//
// func (t *SimpleChaincode) addTutorToTutorial(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	// TODO implement function
// 	// arg0 - tutrorialId (which is the same as tutorialKey), arg1 - tutorId
// 	// Check number of arguments
// 	// Check if tutor and tutorial exist in ledger
// 	// Make sure that tutor is not already listed in tutorial. If that is the case, return error
// 	var tutorialKey, tutorKey string // Entities
//
// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
// 	}
//
// 	tutorialKey = args[0]
// 	tutorKey = args[1]
//
// 	// load tutorial
// 	tutorialB, err := stub.GetState(tutorialKey)
//
// 	if err != nil {
// 		jsonResp := "{\"Error\":\"Failed to get state for " + tutorialKey + "\"}"
// 		return shim.Error(jsonResp)
// 	}
// 	if tutorialB == nil || len(tutorialB) == 0 {
// 		jsonResp := "{\"Error\":\" " + tutorialKey + " does not exit " + "\"}"
// 		return shim.Error(jsonResp)
// 	}
// 	tutorial := tutorial{}
// 	err = json.Unmarshal(tutorialB, &tutorial)
// 	if err != nil {
// 		return shim.Error("Failed to get state")
// 	}
//
// 	// load tutor which will be added to tutorial
// 	tutor, err := stub.GetState(tutorKey)
// 	logger.Info("Tutor " + tutorKey + "postoji")
// 	logger.Info(tutor)
//
// 	if err != nil {
// 		jsonResp := "{\"Error\":\"Failed to get state for " + tutorKey + "\"}"
// 		return shim.Error(jsonResp)
// 	}
// 	if tutor == nil || len(tutor) == 0 {
// 		jsonResp := "{\"Error\":\" " + tutorKey + " does not exit " + "\"}"
// 		return shim.Error(jsonResp)
// 	}
//
// 	for i := 0; i < len(tutorial.Tutors); i++ {
// 		if tutorial.Tutors[i] == tutorKey {
// 			jsonResp := "{\"Error\":\" Tutor with Id: " + tutorKey + " is already on list to tutors" + "\"}"
// 			return shim.Error(jsonResp)
// 		}
// 	}
//
// 	tutorial.Tutors = append(tutorial.Tutors, tutorKey)
//
// 	ajson, _ := json.Marshal(tutorial)
// 	err = stub.PutState(tutorial.Id, ajson)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
//
// 	return shim.Success(nil)
// }
//
// func (t *SimpleChaincode) removeTutorFromTutorial(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	// TODO implement function
// 	// arg0 - tutrorialId, arg1 - tutorId
// 	// Check number of arguments
// 	// Check if tutor and tutorial exist in ledger
// 	// If tutor (which we want to remove) is not listed in tutorial, return error
// 	var tutorialKey, tutorKey string // Entities
//
// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
// 	}
//
// 	tutorialKey = args[0]
// 	tutorKey = args[1]
//
// 	// load tutorial
// 	tutorialB, err := stub.GetState(tutorialKey)
//
// 	if err != nil {
// 		jsonResp := "{\"Error\":\"Failed to get state for " + tutorialKey + "\"}"
// 		return shim.Error(jsonResp)
// 	}
// 	if tutorialB == nil || len(tutorialB) == 0 {
// 		jsonResp := "{\"Error\":\" " + tutorialKey + " does not exit " + "\"}"
// 		return shim.Error(jsonResp)
// 	}
// 	tutorial := tutorial{}
// 	err = json.Unmarshal(tutorialB, &tutorial)
// 	if err != nil {
// 		return shim.Error("Failed to get state")
// 	}
//
// 	// load tutor which will be removed from tutorial
// 	tutor, err := stub.GetState(tutorKey)
// 	logger.Info("Tutor " + tutorKey + "postoji")
// 	logger.Info(tutor)
//
// 	if err != nil {
// 		jsonResp := "{\"Error\":\"Failed to get state for " + tutorKey + "\"}"
// 		return shim.Error(jsonResp)
// 	}
// 	if tutor == nil || len(tutor) == 0 {
// 		jsonResp := "{\"Error\":\" " + tutorKey + " does not exit " + "\"}"
// 		return shim.Error(jsonResp)
// 	}
//
// 	for i := 0; i < len(tutorial.Tutors); i++ {
// 		if tutorial.Tutors[i] == tutorKey {
//
// 			tutorial.Tutors = append(tutorial.Tutors[:i], tutorial.Tutors[i+1:]...)
// 			ajson, _ := json.Marshal(tutorial)
// 			err = stub.PutState(tutorial.Id, ajson)
// 			if err != nil {
// 				return shim.Error(err.Error())
// 			}
//
// 			return shim.Success(nil)
// 		}
// 	}
//
// 	// If tutor is not removed then it does not exits in list of tutors for given tutorial
// 	jsonResp := "{\"Error\":\" Tutor with Id: " + tutorKey + " is not on list of tutors" + "\"}"
// 	return shim.Error(jsonResp)
// }

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	logger.Infof("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}

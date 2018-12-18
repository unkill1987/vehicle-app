package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings" 
	"time"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}


type Vehicle struct {
	Manufacture		string `json:"manufacture"`
	Factory 		string `json:"factory"`
	Name 			string `json:"name"`
	Vehicletype 	string `json:"vehicletype"`
	Volume 			string `json:"volume"`
	Fuel 			string `json:"fuel"`

	Government 		string `json:"government"`
	Plate 			string `json:"plate"`
	Owner 			string `json:"owner"`
	Tradehistory 	string `json:"tradehistory"`
	Price 			string `json:"price"`

	Repair 			string `json:"repair"`
	Repaircosts		string `json:"repaircosts"`
	Shop			string `json:"shop"`

	Accident 		string `json:"accident"`
	Handdlingcosts 	string `json:"handdlingcosts"`
	Insurance 		string `json:"insurance"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()

	if function =="carHistory" {
		return s.carHistory(APIstub, args) 
	}else if function == "Newcar"{
		return s.Newcar(APIstub, args)   
	}else if function == "Change"{
		return s.Change(APIstub, args) 
	}else if function == "Repair"{
		return s.Repair(APIstub, args) 
	}else if function == "Accident"{
		return s.Accident(APIstub, args) 
	}else if function == "Allcars"{
		return s.Allcars(APIstub) 
	}
	
	return shim.Error("Invalid SmartContract function name.")
}

func (s *SmartContract) carHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) !=1{
		return shim.Error("Incorrect number of arguments.Expecting 1")
	}
	sn := args[0]
	history, err := APIstub.GetHistoryForKey(sn)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer history.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for history.HasNext() {
		response, err := history.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
	
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Value\":")
		buffer.WriteString(string(response.Value))
		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("Success")
	
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) Newcar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	
	
	if len(args) != 7  {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	fmt.Println("start init vehicle")
	if len(args[0]) == 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) == 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) == 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) == 0 {
		return shim.Error("4rd argument must be a non-empty string")
	}
	if len(args[4]) == 0 {
		return shim.Error("5rd argument must be a non-empty string")
	}
	if len(args[5]) == 0 {
		return shim.Error("6rd argument must be a non-empty string")
	}
	if len(args[6]) == 0 {
		return shim.Error("7rd argument must be a non-empty string")
	}
	
	sn:= args[0]
	manufacture := strings.ToLower(args[1])
	factory := strings.ToLower(args[2])
	name := strings.ToLower(args[3])
	vehicletype := strings.ToLower(args[4])
	volume := strings.ToLower(args[5])
	fuel := strings.ToLower(args[6])

	idsBytes, err := APIstub.GetState(sn)

	if err != nil {
		return shim.Error(err.Error())
	} else if idsBytes != nil {
		fmt.Println("This vehicle already exists: " + sn)
		return shim.Error("This vehicle already exists: " + sn)
	}
	
	var vehicle = Vehicle{Manufacture:manufacture,Factory:factory,Name:name,Vehicletype:vehicletype,Volume:volume,Fuel:fuel}
	
	vehicleJSONasBytes, err := json.Marshal(vehicle)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = APIstub.PutState(sn,vehicleJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (s *SmartContract) Change(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. you should input 6")
	}

	sn := args[0]
	government := strings.ToLower(args[1])
	plate := strings.ToLower(args[2])
	owner := strings.ToLower(args[3])
	tradehistory := strings.ToLower(args[4])
	price := strings.ToLower(args[5])
	fmt.Println("changes from", sn)

	vehicleAsBytes, err := APIstub.GetState(sn)
	if err != nil {
		return shim.Error("Failed to Record:" + err.Error())
	} else if vehicleAsBytes == nil {
		return shim.Error("Could not found SN")
	}
	vehicleRecord := Vehicle{}
	err = json.Unmarshal(vehicleAsBytes, &vehicleRecord)
	if err != nil{
		return shim.Error(err.Error())
	}
	vehicleRecord.Government = government
	vehicleRecord.Plate = plate
	vehicleRecord.Owner = owner
	vehicleRecord.Tradehistory = tradehistory
	vehicleRecord.Price = price

	vehicleRecord.Repair = ""
	vehicleRecord.Repaircosts = ""
	vehicleRecord.Shop = ""

	vehicleRecord.Accident = ""
	vehicleRecord.Handdlingcosts = ""
	vehicleRecord.Insurance = ""

	vehicleAsBytes, _ = json.Marshal(vehicleRecord)
	err = APIstub.PutState(sn, vehicleAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to input changes: %s", args[0]))
	}
	fmt.Println("Recorded success")
	return shim.Success(nil)
}	

func (s *SmartContract) Repair(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. you should input 4")
	}

	sn := args[0]
	repair := strings.ToLower(args[1])
	repaircosts := strings.ToLower(args[2])
	shop := strings.ToLower(args[3])
	
	fmt.Println("repair from", sn)

	vehicleAsBytes, err := APIstub.GetState(sn)
	if err != nil {
		return shim.Error("Failed to Record:" + err.Error())
	} else if vehicleAsBytes == nil {
		return shim.Error("Could not found SN")
	}
	vehicleRecord := Vehicle{}
	err = json.Unmarshal(vehicleAsBytes, &vehicleRecord)
	if err != nil{
		return shim.Error(err.Error())
	}

	vehicleRecord.Government = ""
	vehicleRecord.Plate = ""
	vehicleRecord.Owner = ""
	vehicleRecord.Tradehistory = ""
	vehicleRecord.Price = ""
	
	vehicleRecord.Repair = repair
	vehicleRecord.Repaircosts = repaircosts
	vehicleRecord.Shop = shop

	vehicleRecord.Accident = ""
	vehicleRecord.Handdlingcosts = ""
	vehicleRecord.Insurance = ""


	vehicleAsBytes, _ = json.Marshal(vehicleRecord)
	err = APIstub.PutState(sn, vehicleAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to input changes: %s", args[0]))
	}
	fmt.Println("Recorded success")
	return shim.Success(nil)
}	

func (s *SmartContract) Accident(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. you should input 4")
	}

	sn := args[0]
	accident := strings.ToLower(args[1])
	handdlingcosts := strings.ToLower(args[2])
	insurance := strings.ToLower(args[3])
	
	fmt.Println("accident from", sn)

	vehicleAsBytes, err := APIstub.GetState(sn)
	if err != nil {
		return shim.Error("Failed to Record:" + err.Error())
	} else if vehicleAsBytes == nil {
		return shim.Error("Could not found SN")
	}
	vehicleRecord := Vehicle{}
	err = json.Unmarshal(vehicleAsBytes, &vehicleRecord)
	if err != nil{
		return shim.Error(err.Error())
	}

	vehicleRecord.Government = ""
	vehicleRecord.Plate = ""
	vehicleRecord.Owner = ""
	vehicleRecord.Tradehistory = ""
	vehicleRecord.Price = ""

	vehicleRecord.Repair = ""
	vehicleRecord.Repaircosts = ""
	vehicleRecord.Shop = ""

	vehicleRecord.Accident = accident
	vehicleRecord.Handdlingcosts = handdlingcosts
	vehicleRecord.Insurance = insurance


	vehicleAsBytes, _ = json.Marshal(vehicleRecord)
	err = APIstub.PutState(sn, vehicleAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to input changes: %s", args[0]))
	}
	fmt.Println("Recorded success")
	return shim.Success(nil)
}	

func (s *SmartContract) Allcars(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "9999999999999999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}


func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

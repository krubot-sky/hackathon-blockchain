package main
//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.
import (
  "encoding/json"
	"fmt"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
  log "github.com/sirupsen/logrus"
)
// SmartContract provides functions for managing a books
type SmartContract struct {
	contractapi.Contract
}
// Book describes basic details of what makes up a book
type Book struct {
  Title       string `json:"title"`
  Author      string `json:"author"`
  Description string `json:"description"`
  ISBN        string `json:"isbn"`
  Owner       string `json:"owner"`
}
// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	books := []Book{
		Book{Title: "The Expanse: Leviathan Wakes", Author: "James S. A. Corey", Description: "As close as you'll get to a hollywood blockbuster in book form.", ISBN: "9781841499895", Owner: "Ben"},
    Book{Title: "The Arrest", Author: "Jonathan Lethem", Description: "An impeccably executed, moving, and wildly inventive tale of madness and narrative at the end of the world. Lethem is at the top of his game.", ISBN: "9781838952167", Owner: "Ben"},
    Book{Title: "The Magic Mountain", Author: "Thomas Mann", Description: "A masterwork, unlike any other... a delight, comic and profound, a new form of language, a new way of seeing.", ISBN: "9780749386429", Owner: "Ben"},
    Book{Title: "The Blind Assassin", Author: "Margaret Atwood", Description: "Atwood has never written with more flair and versatility ... a novel of extraordinary variety and reach ... brilliant.", ISBN: "9781860498800", Owner: "Ben"},
    Book{Title: "Cloud Atlas", Author: "David Mitchell", Description: "Six interlocking lives - one amazing adventure.", ISBN: "9780340822784", Owner: "Ben"},
    Book{Title: "Gravity's rainbow", Author: "Thomas pynchon", Description: "Pychon leaves the rest of the American literary establishment at the starting gate.", ISBN: "9780099533214", Owner: "Ben"},
	}
	for _, book := range books {
		bookAsBytes, _ := json.Marshal(book)
		err := ctx.GetStub().PutState(book.ISBN, bookAsBytes)
		err_ := ctx.GetStub().PutState(book.Title, bookAsBytes)
		if err != nil {
			log.Error("Failed to put to world state. %s", err.Error())
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
		if err_ != nil {
			log.Error("Failed to put to world state. %s", err_.Error())
			return fmt.Errorf("Failed to put to world state. %s", err_.Error())
		}
	}
	return nil
}
// QueryBook returns the book stored in the world state with given isbn
func (s *SmartContract) QueryBook(ctx contractapi.TransactionContextInterface, bookISBN string) (*Book, error) {
	bookAsBytes, err := ctx.GetStub().GetState(bookISBN)
	if err != nil {
		log.Error("Failed to read from world state. %s", err.Error())
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}
	if bookAsBytes == nil {
		log.Error("ISBN %s does not exist", bookISBN)
		return nil, fmt.Errorf("ISBN %s does not exist", bookISBN)
	}
	book := new(Book)
	_ = json.Unmarshal(bookAsBytes, book)
	return book, nil
}
// QueryBook returns the book stored in the world state with given title
func (s *SmartContract) QueryBookByTitle(ctx contractapi.TransactionContextInterface, title string) (*Book, error) {
	bookAsBytes, err := ctx.GetStub().GetState(title)
	if err != nil {
		log.Error("Failed to read from world state. %s", err.Error())
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}
	if bookAsBytes == nil {
		log.Error("ISBN %s does not exist", title)
		return nil, fmt.Errorf("ISBN %s does not exist", title)
	}
	book := new(Book)
	_ = json.Unmarshal(bookAsBytes, book)
	return book, nil
}
// ChangeBookOwner updates the owner field of book with given id in world state
func (s *SmartContract) ChangeBookOwner(ctx contractapi.TransactionContextInterface, bookISBN string, newOwner string) error {
	book, err := s.QueryBook(ctx, bookISBN)
	if err != nil {
		log.Error("No book found with the following ISBN: %s", bookISBN)
		return fmt.Errorf("No book found: %s", err.Error())
	}
	book.Owner = newOwner
	bookAsBytes, _ := json.Marshal(book)
	return ctx.GetStub().PutState(bookISBN, bookAsBytes)
}
// QueryBook returns the book stored in the world state with given isbn
func (s *SmartContract) DeleteBook(ctx contractapi.TransactionContextInterface, bookISBN string) error {
	exists, err := s.AssetExists(ctx, bookISBN)
	if err != nil {
	  return err
	}
	if !exists {
	  return fmt.Errorf("the asset %s does not exist", bookISBN)
	}
	return ctx.GetStub().DelState(bookISBN)
}
// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
	  return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return assetJSON != nil, nil
  }
// QueryBook returns the book stored in the world state with given isbn
func (s *SmartContract) AddBook(ctx contractapi.TransactionContextInterface, bookISBN string, description string, title string,  author string, owner string ) error {
	exists, err := s.AssetExists(ctx, bookISBN)
	if err != nil {
	  return err
	}
	if exists {
	  return fmt.Errorf("the asset %s already exists", id)
	}
	book := Book{
		Title: title, 
		Author: author, 
		Description: description, 
		ISBN: bookISBN, 
		Owner: owner,
   }
	bookAsBytes, err = json.Marshal(book)
	if err != nil {
		log.Error("Failed to marshal book. %s", err.Error())
		return fmt.Errorf("Failed to marshal book. %s", err.Error())
	}
	return ctx.GetStub().PutState(book.ISBN, bookAsBytes)
}
func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error create book store chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting book store chaincode: %s", err.Error())
	}
}
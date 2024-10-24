package types

//S should be capitalized so that it can be exported
type Student struct {
	ID        int    
	FirstName string 
	LastName  string 
	Age       int    
}
//example in json data
// {
// 	"ID": 1,
// 	"FirstName": "John",
// 	"LastName": "Doe",
// 	"Age": 25
// }
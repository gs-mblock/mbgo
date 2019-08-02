package utils
import (
	"testing"
	"github.com/fatih/structs"
)

func TestStructs(t *testing.T){
	type Server struct {
		Name        string `json:"name,omitempty"`
		ID          int
		Enabled     bool
		users       []string // not exported
		//http.Server          // embedded
	}
	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}

	// Convert a struct to a map[string]interface{}
// => {"Name":"gopher", "ID":123456, "Enabled":true}
m := structs.Map(server)
println("m:",m["Name"])

// Convert the values of a struct to a []interface{}
// => ["gopher", 123456, true]
v := structs.Values(server)
println("v:",v)

// Convert the names of a struct to a []string
// (see "Names methods" for more info about fields)
n := structs.Names(server)
println("n:",n)

// Convert the values of a struct to a []*Field
// (see "Field methods" for more info about fields)
f := structs.Fields(server)
println("f:",f)

// Return the struct name => "Server"
//n := structs.Name(server)
//println("n:",n)
// Check if any field of a struct is initialized or not.
h := structs.HasZero(server)
println("h:",h)
// Check if all fields of a struct is initialized or not.
z := structs.IsZero(server)
println("z:",z)
// Check if server is a struct or a pointer to struct
i := structs.IsStruct(server)
println("i:",i)

m["newX1"] ="xx1"
println("x1:",m["newX1"].(string))
m["newX2"] =123
println("x2:",m["newX2"].(int))
}
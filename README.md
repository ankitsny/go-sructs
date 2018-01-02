# Structs and maps demo

> To Run `go run main.go maps-eg.go` 

## Declaration of a struct
```go
    type person struct{
            f_name string
            l_prop2 int
            ...
            propn int
        }

```  

## Methods to initialize a struct
1. `user1 := person{"Vir", "li", ..., propVal}}`

2. Method #2 
    ```go
    user2 := person{
		firstName: "John",
		lastName:  "Doe",
	}
    ```
3. Method #3
    ```go
    var user3 person
	// fmt.Printf("%+v", user3) // display props name
	user3.firstName = "FirstNAME"
	user3.lastName = "LastNAME"
	user3.contact.email = "ankso@ankso.com" // if it has nested strct
	user3.contact.zipCode = 12322
    
    ```



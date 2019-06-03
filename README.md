# Go Manual

## Table of Contents

- Data Types
  - Basic
    - Numbers
    - Strings
    - Booleans
  - Aggregate
    - Arrays
    - Structs
  - Reference
    - Pointers
    - Slices
    - Maps
    - Functions
    - Channels
  - Interface
    - Interfaces
- Go Modules

## Numbers

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

ในภาษา go type int จะเป็น 2's complement arithmetic.
* หมายเหตุ
1's complement ในการหาค่าลบ จะสลับบิต 0 เป็น 1 และสลับบิต 1 เป็น 0
ปัญหาของ 1's complement คือค่า -0 ไม่มีในทางคณิตศาสตร์
2's complement เป็นการนำ 1's complement มา + 1 เพื่อนเลื่อนบิตไปทางขวาทำให้ไม่มีค่า -0

ตัวอย่างการแปลง 1's complement ไปเป็น 2's complement
เลขฐาน 2 จาก 0111 แปลงเป็น 1's complement จะได้ 1000 (เปลี่ยนบิต 1 เป็น 0, เปลี่ยนบิต 0 เป็น 1)
แล้ว 1000 1's complement แปลงเป็น 2's complement จะได้ 1001 (1'complement + 1)

ตัวอย่างการหาค่า 1 และ -1
0000 0001 // เลขฐาน 2 มีค่าเท่ากับ 1
1111 1110 // เลขฐาน 2 แบบ 1's Complement มีค่าเท่ากับ -1
1111 1111 // เลขฐาน 2 แบบ 2's Complement มีค่าเท่ากับ -1 (1's Complement + 1)

ตัวอย่างการหาค่า Min, Max ของ Int และ Uint
fmt.Println(math.MinInt8, math.MaxInt8)
fmt.Println(math.MaxUint8)
```

| Precedence | Operators         |
| ---------- | ----------------- |
| 5          | \* / % << >> & &^ |
| 4          | + - ^ \|          |
| 3          | = != < <= > >=    |
| 2          | &&                |
| 1          | \|\|              |

## Strings

```go
package main

import "fmt"

func main() {
	x := "สวัสดี"
	y := []byte{0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	z := []rune(x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("0x%x,", x[i])
	}

	fmt.Println("")
	fmt.Println(x, len(x))
	fmt.Println(x, len(y))
	fmt.Println(x, len(z))

	fmt.Println("\xe0\xb8\xaa")
	fmt.Println("\xe0\xb8\xa7")
	fmt.Println(string(y[6:9]))
	fmt.Println(string(y[:3]))
	fmt.Printf("%q", z[4])
	fmt.Printf("%q", z[5])

}

```

## Json

```go
package main

import (
	"encoding/json"
	"fmt"
)

var data = `[{"id": 1, "name": "Yo", "age": 36, "completed": false},{"id": 2, "name": "Yea", "age": 32}]`

type Person struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Address   string `json:"address,omitempty"`
	Completed *bool  `json:"completed,omitempty"`
}

func main() {
	p := []Person{}

	json.Unmarshal([]byte(data), &p)

	result, err := json.Marshal(p)

	if err != nil {
		return
	}

	fmt.Println(string(result))
}

```

## การใช้งาน JSON ร่วมกับ ioutil, http

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type Users []struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	users := Users{}
	json.Unmarshal(body, &users)
	users[0].Name = "Yuttasak Pannawat"
	result, err := json.Marshal(users)
	if err != nil {
		return
	}
	fmt.Println(string(result))
}

```

## การใช้งาน JSON NewDecoder, JSON NewEncoder

```go
users := Users{}
	jsonDecode := json.NewDecoder(resp.Body)
	jsonDecode.Decode(&users)
	resp.Body.Close()

	users[0].Name = "Yuttasak Pannawat"

	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.Encode(users)
```

## Error Propagation

```go
package main

import (
	"errors"
	"fmt"
	"log"
)

const dbConnection = true
const balance = 500

func getBalance() (int, error) {
	if !dbConnection {
		return 0, errors.New("getBalance: database is down")
	}

	return balance, nil
}

func withdraw(a int) (int, error) {
	b, err := getBalance()

	if err != nil {
		return 0, fmt.Errorf("withdraw: %v", err)
	}

	if a > b {
		return 0, errors.New("withdraw: insufficient fund")
	}

	return (b - a), nil
}

func main() {
	// log.SetFlags(0)
	amount, err := withdraw(200)

	if err != nil {
		log.Fatalf("main: %v", err)
	}

	fmt.Println("Please collect your money: ", amount)
}

```

## Go Modules

```bash
// init
go mod init github.com/yuttasakcom/go-manual

// show all modules
go list -m all

// get module with version
go get github.com/yuttasakcom/go-manual@v0.1.0

// delete old version
go mod tidy

// check versions
go list -m -version github.com/yuttasakcom/go-manual
```

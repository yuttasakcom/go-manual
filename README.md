# คู่มือการเรียนรู้ Golang ฉบับผู้เริ่มต้น

> สามารถอ่านเอกสารเพิ่มเติมได้ที่ <a href="https://golang.org/doc/" target="_blank">https://golang.org/doc/</a>

## Table of Contents

- Getting Started

  - [Introduce](#introduce)
  - [Why](#why)
  - [Installing](#installing)
    - [Ubuntu](#ubuntu)
    - [MacOS](#macos)

- Golang Basics

  - [Types](#types)
    - [Basic]
      - [Numbers](#numbers)
      - [Strings](#strings)
        - [byte]
        - [character]
        - [rune]
        - [unicode]
        - [utf-8]
        - [string literal]
      - [Booleans](#booleans)
    - [Aggregate]
      - [Arrays]
      - [Structs]
    - [Referance]
      - [Pointers]
      - [Slices]
      - [Maps]
      - [Functions]
      - [Channels]
    - [Interface]
      - [Interfaces]
  - [Variables](#variables)
    - How to Name a Variable
    - Scope
    - Constants
    - Defining Multiple Variables
    - An Example Program
  - [Control Structures]
    - For
    - If
    - Switch
  - Arrays, Slices and Maps
    - Arrays
      - len
    - Slices
      - make
    - Maps
  - Functions
    - Your Second Function
    - Returning Multiple Values
    - Variadic Functions
    - Closure
    - Recursion
    - Defer, Panic & Recover
  - Pointers
    - The \* and & operators
    - new
  - Structs and Interfaces
    - Structs
    - Methods
    - Interfaces
  - Concurrency
    - Goroutines
    - Channels
  - Packages
    - Creating Packages
    - Documentation
  - Testing
  - The Core Packages
    - Strings
    - Input / Output
    - Files & Folders
    - Errors
    - Containers & Sort
    - Hashes & Cryptography
    - Servers
    - Parsing Command Line Arguments
    - Synchronization Primitives

- Books
  - [An introduction to programming in Go](https://www.golang-book.com/books/intro)

## Introduce

> Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

## Why

- Fast
  > เร็วส์
- Learning curve
  > การเรียนรู้อยู่ในระดับปานกลาง ไม่ยาก ไม่ง่าย
- Open Source
  > โอเพ่นซอร์ส
- Designed for Modern Hardware
  > เป็นภาษาที่ถูกออกแบบมาให้ทำงานได้ดีกับฮาร์ดแวร์สมัยใหม่
- Go supports concurrency out of the box (goroutines, channels, select)
  > แยกโปรเซสการทำงานได้หลายโปรเซส เพียงแค่พิมพ์คำว่า "go" หน้าฟังก์ชั่นที่ต้องการ
- Go produces executable binaries that are native to your O.S
  > แปลงซอร์สโค้ดเป็น native ไบนารี่
- Your program becomes deployable the moment it builds, no VMs required
  > โปรแกรมสามารถใช้งานได้ทันทีที่สร้างขึ้น โดยไม่จำเป็นต้องมีเครื่องจำลองการทำงาน
- Go is truly cross platform
  > สามารถทำงานข้ามแพลตฟอร์ม
- Go is garbage collected
  > Go จัดการทรัพยากร(ขยะ)ที่ไม่ถูกใช้งานในระบบให้
- Go has an awesome standard library
  > Go มีไลบารี่มาตราฐานรองรับการทำงานที่หลากหลาย

## Installing

### Ubuntu

#### Setup

```bash
$ cd Downloads && wget https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.11.2.linux-amd64.tar.gz
$ sudo ln -s /usr/local/go/bin/go /usr/bin/go
```

- ตรวจสอบเวอร์ชั่น

```bash
  go version
```

#### Set path

- สร้างโฟลเดอร์ go สำหรับเก็บโปรเจ็กต์

```bash
  mkdir -p go
```

- สร้างโฟลเดอร์ src, pkg, bin ใน โฟลเดอร์ go

```bash
  mkdir src pkg bin
```

- แก้ไขไฟล์ .bashrc ที่ directory Home พิมพ์คำสั่ง

```bash
  # open editor
  vi .bashrc

  # append variable
  export GOPATH=$HOME/go
  export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

  # update .bashrc
  source .bashrc
```

- ตรวจสอบ Environment

```bash
  go env # ถ้า env path ยังไม่เปลี่ยนให้ restart เครื่อง
```

### Numbers

```
int เป็น 2's Complement
2's Complement หาได้จาก (1's Complement + 1)

0000 0001 // เลขฐาน 2 มีค่าเท่ากับ 1
1111 1110 // เลขฐาน 2 แบบ 1's Complement มีค่าเท่ากับ -1
1111 1111 // เลขฐาน 2 แบบ 2's Complement มีค่าเท่ากับ -1 หาได้จากเอา (1's Complement + 1)
```

```go
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

### Strings

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

### Json

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
	v := &p

	json.Unmarshal([]byte(data), v)

	result, err := json.Marshal(p)

	if err != nil {
		return
	}

	fmt.Println(string(result))
}

```

### json, ioutil, http

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

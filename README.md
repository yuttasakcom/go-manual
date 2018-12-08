# คู่มือการเรียนรู้ Golang ฉบับผู้เริ่มต้น

> สามารถอ่านเอกสารเพิ่มเติมได้ที่ <a href="https://golang.org/doc/" target="_blank">https://golang.org/doc/</a>

## Table of Contents

- Getting Started

  - [Introduce](#introduce)
  - [Why](#why)
  - [Installing](#installing)
    - [Ubuntu](#ubuntu)

- Golang Basics

  - [Types](#types)
    - [Numbers](#numbers)
    - [Strings](#strings)
    - [Booleans](#booleans)
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

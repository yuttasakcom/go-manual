## Go 101
Go Programming for beginners
> การเขียนโปรแกรมด้วยภาษา Go (สำหรับผู้เริ่มต้น)

## Table of contents
- Section 1 Introduction
  - [Why Go](#why-go)
  - [Installing Go](#install-go)
    - [Mac](#mac)
    - [Windows](#windows)
    - [Linux Ubuntu](#ubuntu)
  - How to Write Go Code?
    - [Workspace](#workspace)
    - [Hello World](#hello-world)
    - [Let's Build a Rest API Client in Go!](#build)
- Section 2 Key Building Blocks
  - Packages
  - Variables
  - Functions
  - Flow Control
  - Applying Your New Knowledge-Binary Search Algorithm
  - Structs
  - Slices
  - Map
  - Applying Your New Knowledge - Let's Create a Set
## Why Go
- Go is a powerful statically typed programming language, that is both simple and fun<br>
  > Go เป็นภาษาการเขียนโปรแกรมที่มีประสิทธิภาพแบบสแตติกซึ่งง่ายและสนุก
- Open Source — https://github.com/golang/go<br>
  > เป็นโครงการโอเพ่นซอร์สสามารถนำไปใช้งาน หรือพัฒนาต่อยอดได้โดยไม่มีค่าใช้จ่าย ลิงค์ซอร์สโค้ด https://github.com/golang/go
- Fast — Learn, Dev, Compile, Deploy, Run<br>
  > ใช้เวลาเรียนรู้เรื่องไวยากรณ์(syntext)ช่วงแรก หลังจากนั้นทุกอย่างจะเร็วขึ้น Dev, Compile, Deploy, Run
- Designed for Modern Hardware<br>
  > เป็นภาษาที่ถูกออกแบบมาให้ทำงานได้ดีกับฮาร์ดแวร์สมัยใหม่
- Go supports concurrency out of the box (goroutines, channels, select)<br>
  Just add the word "go" and your function runs concurrently<br>
  > รองรับ "การร้องขอเข้าใช้งานทรัพยากรพร้อมกัน" แบบแยกโปรเซสการทำงานได้หลายโปรเซส เพียงแค่พิมพ์คำว่า "go" หน้าฟังก์ชั่นที่ต้องการรันแบบควบคู่กันไป
- Go produces executable binaries that are native to your O.S<br>
  > Go แปลซอร์สโค้ดเป็น native ไบนารี่ให้ทำงานอยู่บนระบบฏิบัติการของเรา
- Your program becomes deployable the moment it builds, no VMs required<br>
  > โปรแกรมสามารถใช้งานได้ทันทีที่สร้างขึ้น โดยไม่จำเป็นต้องมีเครื่องจำลองการทำงาน
- Go is truly cross platform
  > Go เป็นการทำงานแบบข้ามแพลตฟอร์มอย่างแท้จริง
- Write code once, deploy anywhere
  > เขียนโค้ดเพียงครั้งเดียวใช้งานได้ทุกที่
- Go is garbage collected
  > Go จัดการทรัพยากร(ขยะ)ที่ไม่ถูกใช้งานในระบบให้
- Go has an awesome standard library
  > Go มีไลบารี่มาตราฐานรองรับการทำงานที่หลากหลาย
- From web services to encryptions to serialization formats, all what you need to build modern software, all open source
  > ตั้งแต่การทำเว็บเซอร์วิสไปจนถึงรูปแบบการเข้ารหัสต่างๆ ทั้งหมดที่คุณต้องการถูกสร้างด้วยแนวคิดการสร้างซอฟต์แวร์สมัยใหม่ และเป็นโอเพ่นซอร์ส

## Install Go
[Link download](https://golang.org/dl/)<br>
## Mac
Mac — $ brew install go
## Windows
Windows — C:\> choco install golang
## Ubuntu
 * Setup
   * ดาวน์โหลด go ตามลิงค์ [https://golang.org/dl/](https://golang.org/dl/)
   * เข้าไปที่โฟลเดอร์ Download แล้วพิมพ์คำสั่ง
     `sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz`
   * สร้างซิมลิงค์สำหรับเรียก go พิมพ์คำสั่ง
     `sudo ln -s /usr/local/go/bin/go /usr/bin/go`
   * ตรวจสอบเวอร์ชั่น `go version`

 * Set path
   * สร้างโฟลเดอร์ Workspace และโฟลเดอร์ go สำหรับเก็บโปรเจ็กต์
     `mkdir -p Workspace/Back-End/go`
   * สร้างโฟลเดอร์ src, pkg, bin ใน Workspace/Back-End/go<br>
    `mkdir Workspace/Back-End/go && cd Workspace/Back-End/go && mkdir src pkg bin` 
   * แก้ไขไฟล์ .bashrc ที่ directory Home พิมพ์คำสั่ง
     `vi .bashrc`<br>
     export GOPATH=$HOME/Workspace/Back-End/go<br>
     export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
   * ตรวจสอบ Environment `go env`
   * ถ้า env ไม่เปลี่ยนให้ restart เครื่อง
## Workspaces
- src folder: where all your source code will reside
  > โปรเจ็กต์ซอร์สโค้ดทั้งหมดจะถูกเก็บไว้ที่นี่
  <br>ตัวอย่าง เช่น
  ```
  src/
    github.com/yuttasakcom/go101/
  ```
- pkg folder: destination folder for package objects
- bin folder: where your built executables will reside
## Build
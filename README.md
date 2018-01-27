# Go Manual

> https://golang.org/doc/

## Table of Contents

* Section 1

  * [Introduce](#introduce)
  * [Why](#why)
  * [Installing](#installing)
    * [Ubuntu](#ubuntu)

* Section 2

  * [ทำความเข้าใจเรื่อง Package](#package)
  * [ตัวแปร](#variables)
  * Functions
  * Flow Control
  * Applying Your New Knowledge-Binary Search Algorithm
  * Structs
  * Slices
  * Map
  * Applying Your New Knowledge - Let's Create a Set

## Introduce

> Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

## Why

* Fast
  > เร็วส์
* Learning curve
  > การเรียนรู้อยู่ในระดับปานกลาง ไม่ยาก ไม่ง่าย
* Open Source
  > โอเพ่นซอร์ส
* Designed for Modern Hardware
  > เป็นภาษาที่ถูกออกแบบมาให้ทำงานได้ดีกับฮาร์ดแวร์สมัยใหม่
* Go supports concurrency out of the box (goroutines, channels, select)
  > แยกโปรเซสการทำงานได้หลายโปรเซส เพียงแค่พิมพ์คำว่า "go" หน้าฟังก์ชั่นที่ต้องการ
* Go produces executable binaries that are native to your O.S
  > แปลงซอร์สโค้ดเป็น native ไบนารี่
* Your program becomes deployable the moment it builds, no VMs required
  > โปรแกรมสามารถใช้งานได้ทันทีที่สร้างขึ้น โดยไม่จำเป็นต้องมีเครื่องจำลองการทำงาน
* Go is truly cross platform
  > สามารถทำงานข้ามแพลตฟอร์ม
* Go is garbage collected
  > Go จัดการทรัพยากร(ขยะ)ที่ไม่ถูกใช้งานในระบบให้
* Go has an awesome standard library
  > Go มีไลบารี่มาตราฐานรองรับการทำงานที่หลากหลาย

## Installing

### Ubuntu

#### Setup

* ดาวน์โหลด go ตามลิงค์ [https://golang.org/dl/](https://golang.org/dl/)
* เข้าไปที่โฟลเดอร์ Download แล้วพิมพ์คำสั่ง

```bash
  cd Download && sudo tar -C /usr/local -xzf go1.9.3.linux-amd64.tar.gz
```

* สร้างซิมลิงค์สำหรับเรียก go พิมพ์คำสั่ง

```bash
  sudo ln -s /usr/local/go/bin/go /usr/bin/go
```

* ตรวจสอบเวอร์ชั่น

```bash
  go version
```

#### Set path

* สร้างโฟลเดอร์ go สำหรับเก็บโปรเจ็กต์

```bash
  mkdir -p go
```

* สร้างโฟลเดอร์ src, pkg, bin ใน โฟลเดอร์ go

```bash
  mkdir src pkg bin
```

* แก้ไขไฟล์ .bashrc ที่ directory Home พิมพ์คำสั่ง

```bash
  # open editor
  vi .bashrc

  # append variable
  export GOPATH=$HOME/Workspace/Back-End/go
  export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

  # update .bashrc
  source .bashrc
```

* ตรวจสอบ Environment

```bash
  go env # ถ้า env path ยังไม่เปลี่ยนให้ restart เครื่อง
```

#### Workspaces

```
- bin # where your built executables will reside
- pkg # destination folder for package objects
- src
  |_github.com
    |_yuttasakcom
      |_go-sample
```

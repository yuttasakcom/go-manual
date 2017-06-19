# Go programming on Ubuntu

#### การติดตั้ง
 * ทำการดาวน์โหลด : [https://golang.org/dl/](https://golang.org/dl/)
 * เข้าไปที่โฟลเดอร์ Download แล้วพิมพ์คำสั่ง
   `sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz`
 * สร้างซิมลิงค์เรียก go พิมพ์คำสั่ง
   `sudo ln -s /usr/local/go/bin/go /usr/bin/go`
 * ตรวจสอบเวอร์ชั่น `go version`

#### การเซ็ต Path variables
 * สร้างโฟลเดอร์ Workspace และโฟลเดอร์ go สำหรับเก็บโปรเจ็กต์
   `mkdir -p Workspace/go`
 * แก้ไขไฟล์ .profile ที่ directory Home พิมพ์คำสั่ง
   `vi .profile`
   * export GOPATH=$HOME/Workspace/go
   * export PATH=$HOME/Workspace/go:$PATH
 * ตรวจสอบ Environment `go env`

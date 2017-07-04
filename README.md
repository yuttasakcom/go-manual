## Go 101

## สารบัญ
- บทที่ 1 
  - [Install on Ubuntu](#install_ubuntu)

## Install Ubuntu
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

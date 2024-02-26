# เลือกฐานข้อมูล Docker Image ของ Golang
FROM golang:latest

# กำหนดพาธที่เก็บโค้ดของแอปพลิเคชัน
WORKDIR /app

# คัดลอกไฟล์ go.mod และ go.sum ไปยังพาธที่เก็บโค้ด
COPY go.mod .
COPY go.sum .

# ทำการโหลดแพ็คเกจที่ระบุใน go.mod และ go.sum และติดตั้งลงใน Docker Image
RUN go mod download

# คัดลอกโค้ดของแอปพลิเคชันไปยังพาธที่เก็บโค้ด
COPY . .

# คอมไพล์และสร้างไฟล์ binary ของแอปพลิเคชัน
RUN go build -o main .

# กำหนดคำสั่งที่จะใช้ในการเรียกใช้แอปพลิเคชัน
CMD ["./main"]

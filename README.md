# Go-Pokemon API

RESTful API สำหรับจัดการข้อมูล Pokémon พัฒนาด้วยภาษา Go ตามโครงสร้างของ [MarketByGolang (branch dev)](https://github.com/kittiphong064008-cyber/MarketByGolang/tree/dev)

## Tech Stack
- **Framework:** Fiber v3 (`github.com/gofiber/fiber/v3`)
- **Config Management:** Viper (`github.com/spf13/viper`) ผ่าน `config.yml`
- **Validation:** Validator v10 (`github.com/go-playground/validator/v10`)
- **Database:** PostgreSQL (`github.com/lib/pq`)
- **Container:** Docker & Docker Compose

---

## โครงสร้างโปรเจกต์ (ตามแบบ MarketByGolang branch `dev`)

```text
Go-Pokemon/
├── config/
│   └── config.go                 # ใช้ Viper อ่าน config.yml
├── config.yml                    # กำหนดค่า port, db credentials
├── databases/
│   ├── postgres.go               # เชื่อมต่อ PostgreSQL
│   └── schema.sql                # SQL Schema ครบทั้ง 10 ตาราง
├── pkg/
│   └── pokemons/
│       ├── domain/               # model.go, query.go
│       ├── dto/                  # request.go, response.go (ResponseOne, ResponseAll, ResponseRows, ResponseError)
│       ├── repository/           # repository.go
│       ├── service/              # service.go
│       └── handler/              # handler.go (ใช้ validator/v10)
├── server/
│   ├── router/
│   │   └── pokemon.go            # แยก Route กลุ่ม Pokemon
│   └── server.go                 # Server Singleton (sync.Once) + Start()
├── docker-compose.yml
└── main.go
```

---

## การรันเซิร์ฟเวอร์
```bash
go run main.go
```
เซิร์ฟเวอร์จะเริ่มทำงานตามพอร์ตที่กำหนดใน `config.yml` (ค่าเริ่มต้น: `http://localhost:8000`)

---

## API Endpoints

- **GET** `http://localhost:8000/api/v1/pokemon` (ดูรายชื่อ Pokemon ทั้งหมด)
- **GET** `http://localhost:8000/api/v1/pokemon/:id` (ดูรายละเอียดตาม ID)
- **POST** `http://localhost:8000/api/v1/pokemon` (สร้าง Pokemon ใหม่)
- **PATCH** `http://localhost:8000/api/v1/pokemon/:id` (อัปเดตข้อมูล Pokemon บางส่วน)
- **DELETE** `http://localhost:8000/api/v1/pokemon/:id` (ลบ Pokemon ตาม ID)

### ตัวอย่างการ PATCH:
**PATCH** `http://localhost:8000/api/v1/pokemon/25`
```json
{
  "hp": 50,
  "attack": 60
}
```
Response:
```json
{
  "message": "Update Success",
  "RowsAffect": 1
}
```

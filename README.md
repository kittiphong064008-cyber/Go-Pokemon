# Go-Pokemon API

RESTful API สำหรับจัดการข้อมูล Pokémon พัฒนาด้วยภาษา Go ตามหลักการ **Clean Architecture** (Package by Feature)

## Tech Stack
- **Language:** Go 1.27
- **Web Framework:** [Fiber v3](https://github.com/gofiber/fiber/v3)
- **Validation:** [go-playground/validator/v10](https://github.com/go-playground/validator/v10)
- **Database:** PostgreSQL (Driver: `github.com/lib/pq`)
- **Container:** Docker & Docker Compose

---

## โครงสร้างโปรเจกต์ (Clean Architecture)

```text
Go-Pokemon/
├── config/                  # โหลด Environment Variables
├── databases/               # เชื่อมต่อ PostgreSQL และ Auto-migration
├── pkg/
│   └── pokemons/            # Feature: Pokemon
│       ├── domain/          # Entities & Internal Queries (ไม่มี Dependency ภายนอก)
│       ├── dto/             # Data Transfer Objects พร้อม Validation tags
│       ├── repository/      # Layer ติดต่อ Database (SQL queries)
│       ├── service/         # Business Logic
│       └── handler/         # Controller รับ HTTP Request จาก Fiber
├── server/                  # กำหนด Routing และประกอบ Dependencies (DI)
├── .env                     # App & DB Configuration
├── docker-compose.yml       # PostgreSQL & pgAdmin
└── main.go                  # Application Entry Point
```

---

## วิธีการรันโปรเจกต์

### 1. เริ่มต้น Database ผ่าน Docker Compose
```bash
docker compose up -d
```
- **PostgreSQL:** `localhost:5432`
- **pgAdmin:** `http://localhost:5050` (Email: `admin@admin.com`, Pass: `admin`)

### 2. รัน Go Server
```bash
go run main.go
```
เซิร์ฟเวอร์จะเริ่มทำงานที่ `http://localhost:8000`

---

## API Endpoints

### Health Check
- **GET** `http://localhost:8000/health`

### Pokemons
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/api/v1/pokemons` | สร้าง Pokemon ใหม่ |
| `GET` | `/api/v1/pokemons` | ดึงรายชื่อ Pokemon ทั้งหมด |
| `GET` | `/api/v1/pokemons/:id` | ดึงข้อมูล Pokemon ตาม ID |
| `DELETE` | `/api/v1/pokemons/:id` | ลบ Pokemon ตาม ID |

#### ตัวอย่าง Request Body (POST `/api/v1/pokemons`):
```json
{
  "name": "Pikachu",
  "type": "Electric",
  "hp": 100,
  "attack": 55
}
```
*(มีการตรวจสอบข้อมูลด้วย `validator/v10`: name ต้องมี 2-100 ตัวอักษร, hp ระหว่าง 1-1000, attack ระหว่าง 1-500)*

# Go Base API - Arhitektuuri Plaan

## Ülevaade

Go REST API, mis kasutab **Fiber frameworki** ja töötab **Docker** konteineris. API sisaldab **API Key autentimist** ja järgib puhta arhitektuuri põhimõtteid.

---

## Projekti Kaustastruktuur

```
api/
├── cmd/
│   └── server/
│       └── main.go              # Entry point - käivitab serveri
├── internal/
│   ├── config/
│   │   └── config.go            # Konfiguratsiooni laadimine .env failist
│   ├── middleware/
│   │   └── auth.go              # API Key autentimise middleware
│   ├── routes/
│   │   └── routes.go            # Kõik route definitsioonid
│   ├── controllers/
│   │   ├── health_controller.go # Health check endpoint
│   │   └── user_controller.go   # Kasutajate CRUD controller
│   ├── services/
│   │   └── user_service.go      # Kasutajate ärispetsiifiline loogika
│   ├── models/
│   │   └── user.go              # Andmemudelid (structid)
│   └── utils/
│       ├── response.go          # API response helperid
│       └── validator.go         # Valideerimise utiliidid
├── .env                         # Keskkonnamuutujad (API key jne)
├── .env.example                 # Näidis .env fail
├── .dockerignore                # Docker build ignore reeglid
├── Dockerfile                   # Docker image definitsioon
├── docker-compose.yml           # Docker Compose konfiguratsioon
├── go.mod                       # Go module definitsioon
├── go.sum                       # Dependency checksumid
└── README.md                    # Dokumentatsioon
```

---

## Arhitektuuri Kirjeldus

### 1. **cmd/server/main.go** - Entry Point
- Laeb konfiguratsiooni
- Initsialiseerib Fiber appi
- Registreerib middleware'id
- Registreerib route'id
- Käivitab serveri määratud portil

### 2. **internal/config/config.go** - Konfiguratsioon
- Kasutab `godotenv` .env faili laadimiseks
- Struktuur konfiguratsiooni hoidmiseks:
  - `ServerPort` - serveri port (vaikimisi 3000)
  - `APIKey` - autentimise API key
  - `Environment` - keskkond (development/production)

### 3. **internal/middleware/auth.go** - API Key Autentimine
- Kontrollib iga requesti `X-API-Key` headerit
- Võrdleb headeris olevat key'd konfiguratsioonis olevaga
- Tagastab 401 Unauthorized, kui key puudub või on vale
- Health check endpoint jäetakse autentimisest välja

### 4. **internal/routes/routes.go** - Route Definitsioonid
- Registreerib kõik endpointid
- Rakendab autentimise middleware'i kaitstud route'idele
- Route struktuur:
  ```
  GET  /api/health          - Health check (ilma autentimiseta)
  GET  /api/users           - Kõik kasutajad
  GET  /api/users/:id       - Üks kasutaja
  POST /api/users           - Loo kasutaja
  PUT  /api/users/:id       - Uuenda kasutajat
  DELETE /api/users/:id     - Kustuta kasutaja
  ```

### 5. **internal/controllers/** - Controllerid
- **health_controller.go**: Tagastab API staatuse
- **user_controller.go**: CRUD operatsioonid kasutajate jaoks
  - Kutsub välja service kihti ärispetsiifilise loogika jaoks
  - Kasutab utils'e response'ide jaoks

### 6. **internal/services/user_service.go** - Ärispetsiifiline Loogika
- Sisaldab kasutajate loogikat (hetkel in-memory andmed)
- Tulevikus lihtne asendada andmebaasi ühendusega
- Funktsioonid:
  - `GetAllUsers()`
  - `GetUserByID(id)`
  - `CreateUser(user)`
  - `UpdateUser(id, user)`
  - `DeleteUser(id)`

### 7. **internal/models/user.go** - Andmemudelid
- `User` struct koos JSON tag'idega
- `CreateUserRequest` - valideerimiseks
- `UpdateUserRequest` - valideerimiseks

### 8. **internal/utils/** - Utility Funktsioonid
- **response.go**: Standardiseeritud API response'id
  - `SuccessResponse(data)`
  - `ErrorResponse(message, statusCode)`
  - `NotFoundResponse(message)`
- **validator.go**: Sisendi valideerimine
  - Email valideerimine
  - Required field kontroll

---

## API Key Autentimise Flow

```
Client Request
    │
    ▼
┌─────────────────┐
│  Fiber Server   │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Auth Middleware │
│  X-API-Key?     │
└────────┬────────┘
         │
    ┌────┴────┐
    │         │
    ▼         ▼
  [Ei]      [Jah]
    │         │
    ▼         ▼
  401      Kontrolli
  Error    key õigsust
              │
         ┌────┴────┐
         │         │
         ▼         ▼
       [Vale]    [Õige]
         │         │
         ▼         ▼
       401      Controller
       Error
```

---

## Docker Konfiguratsioon

### Dockerfile
- Multi-stage build (builder + runtime)
- Alpine base image väiksema suuruse jaoks
- Non-root user turvalisuse jaoks

### docker-compose.yml
- API teenus
- Port mapping
- Environment variables
- Volume mount .env faili jaoks

---

## Näidis .env Fail

```env
SERVER_PORT=3000
API_KEY=your-secret-api-key-here
ENVIRONMENT=development
```

---

## API Response Formaat

### Edukas Response
```json
{
  "success": true,
  "data": { ... },
  "message": "Operatsioon õnnestus"
}
```

### Viga Response
```json
{
  "success": false,
  "error": "Viga kirjeldus",
  "code": 400
}
```

---

## Järgmised Sammad

1. Loo kaustastruktuur
2. Initsialiseeri Go module
3. Paigalda Fiber ja godotenv
4. Loo konfiguratsiooni haldus
5. Loo middleware
6. Loo models, services, controllers
7. Loo utils
8. Loo routes
9. Loo main.go
10. Loo Docker failid
11. Loo README (arhitektuur, projekti setup ja printsiibid)
12. Testi API-d


# Go Base API

Go REST API, mis kasutab [Fiber](https://gofiber.io/) frameworki ja töötab Docker konteineris. API sisaldab API Key autentimist ja järgib puhta arhitektuuri põhimõtteid.

## 📋 Sisukord

- [Funktsioonid](#funktsioonid)
- [Arhitektuur](#arhitektuur)
- [Kiire Start](#kiire-start)
- [API Endpoints](#api-endpoints)
- [Autentimine](#autentimine)
- [Konfiguratsioon](#konfiguratsioon)
- [Docker](#docker)
- [Arendus](#arendus)

## ✨ Funktsioonid

- ✅ **Fiber Framework** - Kiire ja minimalistlik Go web framework
- ✅ **API Key Autentimine** - Turvaline autentimine X-API-Key headeriga
- ✅ **Puhas Arhitektuur** - Eraldi kihid: routes, controllers, services, models
- ✅ **Docker Tugi** - Multi-stage Dockerfile ja docker-compose
- ✅ **Environment Konfiguratsioon** - .env faili tugi
- ✅ **Standardiseeritud Response'id** - Ühtne JSON response formaat
- ✅ **Health Check** - API staatuse kontrollimise endpoint
- ✅ **CRUD Operatsioonid** - Näidis kasutajate haldusega

## 🏗️ Arhitektuur

```
api/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Konfiguratsiooni laadimine
│   ├── middleware/
│   │   └── auth.go              # API Key autentimine
│   ├── routes/
│   │   └── routes.go            # Route definitsioonid
│   ├── controllers/
│   │   ├── health_controller.go # Health check
│   │   └── user_controller.go   # Kasutajate CRUD
│   ├── services/
│   │   └── user_service.go      # Ärispetsiifiline loogika
│   ├── models/
│   │   └── user.go              # Andmemudelid
│   └── utils/
│       ├── response.go          # Response helperid
│       └── validator.go         # Valideerimine
├── .env                         # Keskkonnamuutujad
├── .env.example                 # Näidis .env
├── Dockerfile                   # Docker image
├── docker-compose.yml           # Docker Compose
├── go.mod                       # Go module
└── README.md                    # Dokumentatsioon
```

### Arhitektuuri Kihid

| Kiht | Kaust | Kirjeldus |
|------|-------|-----------|
| **Entry Point** | `cmd/server/` | Serveri käivitamine |
| **Config** | `internal/config/` | Konfiguratsiooni haldus |
| **Middleware** | `internal/middleware/` | Autentimine, logimine |
| **Routes** | `internal/routes/` | URL marsruutimine |
| **Controllers** | `internal/controllers/** | HTTP requestide töötlemine |
| **Services** | `internal/services/` | Ärispetsiifiline loogika |
| **Models** | `internal/models/` | Andmestruktuurid |
| **Utils** | `internal/utils/` | Abifunktsioonid |

## 🚀 Kiire Start

### Eeltingimused

- [Go 1.21+](https://golang.org/dl/) (kohalikuks arenduseks)
- [Docker](https://www.docker.com/) (konteineris käitamiseks)
- [Docker Compose](https://docs.docker.com/compose/)

### Kohalik Arendus

1. **Klooni projekt**
   ```bash
   cd api
   ```

2. **Kopeeri .env fail**
   ```bash
   cp .env.example .env
   ```

3. **Muuda .env faili** - asenda API_KEY oma väärtusega

4. **Paigalda sõltuvused**
   ```bash
   go mod tidy
   ```

5. **Käivita server**
   ```bash
   go run cmd/server/main.go
   ```

Server käivitub aadressil `http://localhost:3000`

### Docker

1. **Käivita Docker Compose'ga**
   ```bash
   docker-compose up -d
   ```

2. **Kontrolli staatust**
   ```bash
   docker-compose ps
   ```

3. **Vaata logisid**
   ```bash
   docker-compose logs -f api
   ```

4. **Peata konteinerid**
   ```bash
   docker-compose down
   ```

## 🔗 API Endpoints

### Health Check

| Method | Path | Autentimine | Kirjeldus |
|--------|------|-------------|-----------|
| GET | `/api/health` | ❌ Ei | API staatuse kontroll |

**Response:**
```json
{
  "success": true,
  "data": {
    "status": "healthy",
    "version": "1.0.0"
  },
  "message": "API is running"
}
```

### Kasutajad (Users)

| Method | Path | Autentimine | Kirjeldus |
|--------|------|-------------|-----------|
| GET | `/api/users` | ✅ Jah | Kõik kasutajad |
| GET | `/api/users/:id` | ✅ Jah | Üks kasutaja |
| POST | `/api/users` | ✅ Jah | Loo kasutaja |
| PUT | `/api/users/:id` | ✅ Jah | Uuenda kasutajat |
| DELETE | `/api/users/:id` | ✅ Jah | Kustuta kasutaja |

#### GET /api/users
```bash
curl -H "X-API-Key: your-api-key" http://localhost:3000/api/users
```

#### GET /api/users/:id
```bash
curl -H "X-API-Key: your-api-key" http://localhost:3000/api/users/1
```

#### POST /api/users
```bash
curl -X POST \
  -H "X-API-Key: your-api-key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Johnson", "email": "alice@example.com"}' \
  http://localhost:3000/api/users
```

#### PUT /api/users/:id
```bash
curl -X PUT \
  -H "X-API-Key: your-api-key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Updated"}' \
  http://localhost:3000/api/users/1
```

#### DELETE /api/users/:id
```bash
curl -X DELETE \
  -H "X-API-Key: your-api-key" \
  http://localhost:3000/api/users/1
```

## 🔐 Autentimine

API kasutab **API Key autentimist** läbi `X-API-Key` headeri.

### Kuidas see töötab

1. Iga kaitstud endpointi request peab sisaldama `X-API-Key` headerit
2. Server võrdleb headeris olevat key'd konfiguratsioonis olevaga
3. Kui key puudub või on vale → **401 Unauthorized**
4. Kui key on õige → request jätkub controllerisse

### Erandid

- `/api/health` endpoint ei nõua autentimist

### Näide

```bash
# Õige - 200 OK
curl -H "X-API-Key: your-secret-api-key" http://localhost:3000/api/users

# Vale key - 401 Unauthorized
curl -H "X-API-Key: wrong-key" http://localhost:3000/api/users

# Puudub key - 401 Unauthorized
curl http://localhost:3000/api/users
```

## ⚙️ Konfiguratsioon

Konfiguratsioon toimub `.env` faili kaudu:

| Muutuja | Vaikimisi | Kirjeldus |
|---------|-----------|-----------|
| `SERVER_PORT` | `3000` | Serveri port |
| `API_KEY` | - | **Kohustuslik** - API autentimise võti |
| `ENVIRONMENT` | `development` | Keskkond (development/production) |

### .env Näide

```env
SERVER_PORT=3000
API_KEY=my-super-secret-api-key-12345
ENVIRONMENT=development
```

## 🐳 Docker

### Dockerfile

Kasutab **multi-stage build** optimeerimiseks:
- **Builder stage**: Go kompileerimine
- **Runtime stage**: Alpine Linux, ainult binaarfail

### docker-compose.yml

```yaml
services:
  api:
    build: .
    ports:
      - "3000:3000"
    environment:
      - API_KEY=${API_KEY}
```

### Docker Käsud

```bash
# Build ja käivita
docker-compose up -d --build

# Ainult build
docker build -t go-base-api .

# Käivita üksik konteiner
docker run -p 3000:3000 -e API_KEY=secret go-base-api
```

## 🛠️ Arendus

### Projekti Struktuuri Selgitus

- **`cmd/server/`** - Rakenduse käivitamise punkt
- **`internal/`** - Privaatne kood, mida ei impordita väljastpoolt
- **`config/`** - Keskkonnamuutujate laadimine
- **`middleware/`** - HTTP middleware (autentimine)
- **`routes/`** - URL-i ja controlleri seosed
- **`controllers/`** - HTTP requestide töötlemine
- **`services/`** - Ärireeglid ja loogika
- **`models/`** - Andmestruktuurid
- **`utils/`** - Korduvkasutatavad abifunktsioonid

### Response Formaat

Kõik API response'id järgivad ühtset formaati:

```json
{
  "success": true,
  "data": { ... },
  "message": "Operatsioon õnnestus"
}
```

Vea korral:
```json
{
  "success": false,
  "error": "Vea kirjeldus",
  "code": 400
}
```

### Uue Endpointi Lisamine

1. **Loo model** (`internal/models/`)
2. **Loo service** (`internal/services/`)
3. **Loo controller** (`internal/controllers/`)
4. **Registreeri route** (`internal/routes/routes.go`)

## 📝 Litsents

MIT

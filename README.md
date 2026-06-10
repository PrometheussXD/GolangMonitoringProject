# Website Monitoring System

A lightweight website monitoring system built with Go, Gin, PostgreSQL, and GORM.

The application continuously monitors websites, stores monitoring results, calculates uptime percentages, tracks response times, and provides a web dashboard for visualization and management.

---

## Features

* Monitor multiple websites/endpoints
* Store monitoring history in PostgreSQL
* Measure HTTP response times
* Track website availability (UP/DOWN)
* Calculate uptime percentage
* Dashboard with real-time status overview
* Search monitored targets
* Target details page with monitoring history
* Automatic dashboard refresh
* Telegram alerts for downtime events
* REST API built with Gin
* Database persistence using PostgreSQL and GORM

---

## Tech Stack

### Backend

* Go
* Gin
* GORM

### Database

* PostgreSQL

### Frontend

* HTML
* Bootstrap 5

### Notifications

* Telegram Bot API

---

## Project Structure

```text
cmd/
└── server/

internal/
├── api/
├── dataBase/
├── handlers/
├── models/
├── monitor/
├── repository/
└── telegram/

templates/

.env
```

---

## Database Schema

### Targets

Stores monitored websites.

| Field    | Type   |
| -------- | ------ |
| ID       | uint   |
| Name     | string |
| URL      | string |
| Interval | int    |

### Check Results

Stores monitoring results.

| Field        | Type      |
| ------------ | --------- |
| ID           | uint      |
| TargetID     | uint      |
| StatusCode   | int       |
| ResponseTime | int       |
| IsUp         | bool      |
| CheckedAt    | time.Time |

---

## Environment Variables

Create a `.env` file:

```env
BOT_TOKEN=your_telegram_bot_token
CHAT_ID=your_chat_id
```
## Telegram Setup

### 1. Create a Telegram Bot

Open Telegram and search for:

```text
@BotFather
```

Create a new bot:

```text
/newbot
```

BotFather will return a bot token:

```text
123456789:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

Add this token to your `.env` file:

```env
BOT_TOKEN=your_bot_token
```

---

### 2. Get Your Chat ID

Start a conversation with your bot and send any message:

```text
hello
```

Then open:

```text
https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates
```

Look for:

```json
{
  "message": {
    "chat": {
      "id": 123456789
    }
  }
}
```

The value of `chat.id` is your Telegram Chat ID.

Add it to your `.env` file:

```env
CHAT_ID=123456789
```

---

### 3. Test Notifications

When a monitored website becomes unavailable, the system will automatically send a Telegram notification:

```text
🚨 Google is DOWN
```

---

## Installation

### Clone Repository

```bash
git clone https://github.com/PrometheussXD/GolangMonitoringProject.git
cd website-monitoring-system
```

### Install Dependencies

```bash
go mod tidy
```

### Configure PostgreSQL

Create a database:

```sql
CREATE DATABASE gomonitor;
```

Configure your database connection settings inside the project.

### Run

```bash
go run ./cmd/server
```

Application will start on:

```text
http://localhost:8080
```

Dashboard:

```text
http://localhost:8080/dashboard
```

---

## API Endpoints

### Health Check

```http
GET /health
```

### Create Target

```http
POST /targets
```

Example:

```json
{
  "name": "Google",
  "url": "https://google.com",
  "interval": 30
}
```

### Get Targets

```http
GET /targets
```

### Dashboard

```http
GET /dashboard
```

### Target Details

```http
GET /targets/:id/details
```

### Delete Target

```http
POST /targets/:id/delete
```

---

## Monitoring Workflow

```text
Target
   ↓
HTTP Check
   ↓
Measure Response Time
   ↓
Store Result
   ↓
Update Dashboard
   ↓
Send Telegram Alert (if DOWN)
```

---

## Screenshots
### Dashboard
<img width="1919" height="848" alt="image" src="https://github.com/user-attachments/assets/308ba2dd-311a-4b61-8653-32c56e90e882" />

### Target Details
<img width="1888" height="814" alt="image" src="https://github.com/user-attachments/assets/1a224bb8-15b0-4a36-b16b-563ae8aa22cf" />

### Telegram Bot
<img width="1402" height="815" alt="image" src="https://github.com/user-attachments/assets/8bab7a7f-61ef-4677-aa3e-8d963099da36" />

---

## Future Improvements

* Docker support
* Edit target functionality
* Advanced statistics
* Charts and visual analytics
* Email notifications
* User authentication
* Multi-user support
---

## What I Learned

Through this project I practiced:

* Building REST APIs with Gin
* Database design with PostgreSQL
* Using GORM ORM
* Repository Pattern
* Background monitoring tasks
* Environment variable management
* Telegram Bot API integration
* Server-side rendering with Go templates
* Web dashboard development
---

## Author
Built by Amirreza Radfar as a backend portfolio project while learning Go and software engineering.

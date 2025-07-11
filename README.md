# 🛡️ secSafe

secSafe is a modular SaaS-based physical access control platform. It allows organizations to manage, monitor, and analyze physical access to doors within departments, using QR or NFC-based authentication. Built with Go (Gin), PostgreSQL, Redis, and designed with security, scalability, and resilience in mind.

---

## 📚 Features

- 🔐 Secure user authentication (admin-managed)
- 🧾 Entry & exit logging with timestamps
- 🔄 Door controllers with offline fallback support
- 📈 Security analytics and access summaries
- ⚠️ WhatsApp & Email notifications on access
- 🛂 Multi-factor authentication for admins
- 💬 REST API documented via Swagger
---

## 🧱 Tech Stack

| Layer | Technology |
|-------|------------|
| Backend | Go (Gin) |
| Cache / Queue | Redis |
| Database | PostgreSQL |
| API Docs | Swagger (Swaggo) |
| Notification | WhatsApp (via Twilio/BulkAPI) + SMTP |
| Auth | JWT (Access + Refresh), MFA |
| Deployment | Railway / Docker-ready |
| Monitoring | Zap logger, future: Prometheus + Grafana |

---

| Module       | Description                                         |
| ------------ | --------------------------------------------------- |
| Auth         | Admin-controlled registration, login, JWT & MFA     |
| Door         | Accepts QR/NFC scans, checks access rules           |
| Logs         | Stores access logs, entry/exit analytics            |
| Notification | WhatsApp/email alerts with background job queue     |
| Admin        | Password resets, MFA setup, role management         |
| Cache        | Redis for rate limits, MFA codes, JWT blacklist     |
| Monitoring   | Zap logger (structured), future: Prometheus metrics |


📧 Contact
Author: John Kiribata
Email: johnkiribata12@gmail.com
Project: https://github.com/throneCoder/secSafe



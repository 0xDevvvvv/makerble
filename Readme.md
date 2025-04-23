
## ⚙️ Setup Instructions

1. **Clone the repo**

```bash
git clone https://github.com/0xDevvvvv/makerble
cd makerble
go mod tidy
```

2. **Set up `.env`**

```env
PORT=8080
DBHOST=localhost
DBPORT=5432
DBUSER=youruser
DBPASSWORD=yourpassword
DBNAME=yourdatabasename
JWTSecret=yoursecret
SSL_MODE=disable
```

3. **Run the Server**

```bash
cd cmd
go run main.go
```

---

## 📬 API Endpoints (Example)

| Method | Endpoint            | Role         | Description                        |
|--------|---------------------|--------------|------------------------------------|
| POST   | `/signup`            | All         | signup with username & password    |
| POST   | `/login`            | All          | Login with username & password     |
| POST   | `/patients`         | Receptionist | Register new patient               |
| GET    | `/patients/:id`     | Both         | View patient details               |
| GET    | `/patients/   `     | Both         | View all patient details           |
| PUT    | `/patients/:id`     | Both         | Update diagnosis or notes          |
| DELETE | `/patients/   `     | Receptionist | Delete patient                     |

🧾 Full API Collection: [Postman Collection Link](#) *(link will be pasted soon )*
---

## 🙌 Acknowledgements

This project is part of the internship process at [Makerble](https://makerble.com). Built with ❤️ by [Dev Bhagavan](https://dev-bhagavan.vercel.app).

---

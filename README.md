# 🚀 FastAPI Resume Processing API

Backend API built with **FastAPI** for handling authentication, resume processing, vector database management, and file operations.

---

# Project Requirements
- postgresql
- .env files with api keys for AI, Backend, Frontend Applications


## Python Application

## ⚙️ Requirements

- Python 3.9+
- pip

---

## 📦 Install Dependencies

Move into the Python project directory:

cd AI/application

Install all required packages:

pip install -r requirements.txt

---

## ▶️ Run the Application

uvicorn main:app --reload

---

## 🌐 API Access

After the server is running:

API Base     : http://127.0.0.1:8000  
Swagger UI   : http://127.0.0.1:8000/docs  
ReDoc        : http://127.0.0.1:8000/redoc

---

## 🧪 Test Example

curl http://127.0.0.1:8000

---

## 🧹 Dependency Maintenance (Optional)

After adding or removing packages:

pip freeze > requirements.txt

---

## 📝 Notes

- All resume files are stored in resumes_collection/
- Vector operations and embeddings are handled in vectorDB.py
- Authentication logic is centralized in auth.py

## GO Application

# 🦫 Gin Resume Processing API

High-performance REST API built with **Gin (Go)** for authentication, user management, job handling, middleware processing, and service operations.

---


## ⚙️ Requirements

- Go 1.20+

---

## 📦 Install Dependencies

Move into the Gin project directory:

cd backend

Download and install all required libraries:

go mod tidy

---

## ▶️ Run the Application

go run main.go

---

## 🌐 API Access

After the server is running:

API Base : http://localhost:8080

---

## 🧪 Test Example

curl http://localhost:8080

---

## 🧹 Dependency Maintenance

After adding or removing packages:

go mod tidy

---

## 📝 Notes

- Routing and application startup are handled in main.go
- Business logic is organized under service/
- Request validation and authentication handled in middleware/ and auth/
- All data models are defined inside model/

## vue js (Frontend) application
# 🎨 Frontend Web Application

Frontend application built with **Vue.js**, **Vite**, and **Tailwind CSS**.

--

## ⚙️ Requirements

- Node.js 18+
- npm

---

## 📦 Install Dependencies

Move into the frontend project directory:

cd frontend

Install all required packages:

npm install

---

## ▶️ Run Development Server

npm run dev

---

## 🌐 Access Application

After the server is running:

Local URL : http://localhost:5173

---

## 🧪 Build for Production

npm run build

---

## 👀 Preview Production Build

npm run preview

---

## 📝 Notes

- Vue.js is used as the frontend framework
- Vite is used as the build tool and dev server
- Tailwind CSS handles styling
- All application logic is inside the src/ directory
- Static assets are placed in public/

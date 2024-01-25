# Варіант 13%1 = 1 - Валюти
# LabBack API 

## Початок роботи

Ці інструкції допоможуть вам скопіювати проект та запустити його на вашому локальному комп'ютері для розробки та тестування.

### Передумови

- Go (принаймні версія 1.21.6)
- Docker та Docker Compose

### Встановлення

1. **Клонувати репозиторій**
   ```sh
   git clone git@github.com:LapaJr31/labBackGO.git
   cd labBackGo
   ```

   Збудувати та запустити за допомогою Docker Compose
   ```sh
   docker-compose up --build
   ```

   Це налаштує всю середу, включаючи додаток Go та базу даних PostgreSQL.

   Ініціалізувати базу даних (якщо автоматична міграція не виконується)
   - Підключіться до бази даних PostgreSQL та налаштуйте схему за потребою.

## Маршрути API

API надає наступні кінцеві точки:

### Управління користувачами

- POST /api/user - Створити нового користувача.
- GET /api/user/{id} - Отримати користувача за ID.
- PUT /api/user/{id} - Оновити користувача за ID.
- DELETE /api/user/{id} - Видалити користувача за ID.

### Управління категоріями

- POST /api/category - Створити нову категорію.
- GET /api/category/{id} - Отримати категорію за ID.
- PUT /api/category/{id} - Оновити категорію за ID.
- DELETE /api/category/{id} - Видалити категорію за ID.
- GET /api/categories - Отримати всі категорії.

### Управління записами витрат

- POST /api/expense-record - Створити новий запис витрат.
- GET /api/expense-record/{id} - Отримати запис витрат за ID.
- PUT /api/expense-record/{id} - Оновити запис витрат за ID.
- DELETE /api/expense-record/{id} - Видалити запис витрат за ID.

### Управління валютами

- POST /api/currency - Створити нову валюту.
- GET /api/currency/{id} - Отримати валюту за ID.
- PUT /api/currency/{id} - Оновити валюту за ID.
- DELETE /api/currency/{id} - Видалити валюту за ID.

### Аутентифікація

- POST /api/register - Зареєструвати нового користувача та отримати токен JWT.
- POST /api/login - Увійти з існуючими обліковими даними користувача та отримати токен JWT.
- GET /api/protected - Отримати доступ до захищеного маршруту (потрібен токен JWT).

## Тестування API

Ви можете тестувати кінцеві точки API за допомогою інструментів, таких як Postman або cURL. Переконайтеся, що перед відправкою запитів ви запустили сервер за допомогою Docker Compose.

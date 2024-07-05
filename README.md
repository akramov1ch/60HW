### Uyga Vazifa: Squirrel bilan oddiy `RESTful API` yaratish

#### Maqsad
Ushbu uy vazifasining maqsadi `Go` yordamida `task`larni boshqarish uchun oddiy `RESTful API`, SQL so'rovlarini yaratish uchun `Squirrel` kutubxonasi va ma'lumotlar bazasi sifatida PostgreSQL yaratishdir. 

### Talablar
1. Tasks jadvali bilan PostgreSQL ma'lumotlar bazasini yarating
    - Tasks jadvalida quyidagi ustunlar bo'lishi kerak: `id`, `title`, `description`, `done`.
2. Quyidagi `API` endpointlarni amalga oshiring:
    - `GET /tasks`: Barcha vazifalarni olish.
    - `GET /tasks/{id}`: ID bo'yicha topshiriqni olish.
    - `POST /tasks`: Yangi vazifa yarating.
    - `PUT /tasks/{id}`: ID boʻyicha vazifani yangilang.
    - `DELETE /tasks/{id}`: ID boʻyicha vazifani oʻchirish.
3. PostgreSQL ma'lumotlar bazasi bilan ishlash uchun `SQL` so'rovlarini yaratish uchun `Squirrel` kutubxonasidan foydalaning.


   



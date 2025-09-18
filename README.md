# HTTP-ToDo--Go
Простой ToDo-лист на Go с HTTP API: in‑memory хранилище задач, потокобезопасные операции, JSON DTO и обработка ошибок, роутинг на базе Gorilla Mux.

Эндпоинты
POST /tasks — создать задачу

Body: {"title":"...", "description":"..."}

201 Created, JSON задачи.

GET /tasks/{title} — получить задачу по заголовку

200 OK или 404 если не найдена.

GET /tasks — получить все задачи

200 OK, список JSON.

GET /tasks?completed=true — получить только незавершённые задачи

200 OK, список JSON.

PATCH /tasks/{title} — завершить или снять завершение

Body: {"complete":true|false} 

200 OK, обновлённая задача.

DELETE /tasks/{title} — удалить задачу

204 No Content.

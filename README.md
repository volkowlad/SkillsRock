# SkillsRock

## Выполнено тестовое задание
### В проекте выполнены следующие задачи:
- Выполнены REST методы.
- Работа с фреймворком fiber.
- Автоматически сгенерирован swagger документация с помощью библиотеки swaggo.
- Работа с БД PostgreSQL: выполнены миграции с помощью migrate, запуск, используя Docker и Docker Compose.
- Для работы с БД ипользовалась стандартная библиотека pgx. Выполнение SQL запросов, выполнение транзакций.
- Работа с переменными окружения .env. Создан файл .env.example для примера файла .env.
- Работа с конфигом осуществлена с помощью библиотеки viper.
- Сделан Makefile для удобной работы с проектом
- Graceful Shutdown

### Для запуска приложения:
```
make run
```
### Для остановки приложения:
```
make down
```
### Тестовое задание для Effective Mobile

Запуск локально...
1. docker-compose -f ./deploy/docker-compose.yml --env-file ./configs/.env up -d --remove-orphans --build

2. Ожидаем запуска...

3. goose -dir schema postgres 'postgresql://EffectiveMobile:somestrongpassword@localhost:5432/db_peoples' up
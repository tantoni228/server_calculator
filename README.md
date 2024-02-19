# Distributed calculator
# Описание: 
1) Вычисляет пример, заданный пользователем;
2) Показывает ходы решения примера;

# Графический интерфес: 
не реализован

# Запуск:
1) Скачать **Zip** файл и разархивировать его;
2) Открыть папку в среде разработке (у меня **Vs Code**);
3) В терминале нужно ввести команду, которая установит библиотеки;
```rb
go mod download
```
4) Дальше следует запустить сервер **MySql**;
Я использовал приложение [MAMP ссылка на скачивание](https://www.mamp.info/en/downloads/)
Можно посмотреть видео, как запустить [сервер My Sql](https://www.youtube.com/watch?v=4Wf__mTxm8M)

## Краткое описание
1. Нажмите *Start servers*
2. После нажмите на *Open web start*
3. Вас перебросит на сервер *MAMP*
4. Дальше найдите **phpMyAdmin**, блягодаря нему вы включите сервер My Sql
5. Вы уже зашли и введите панель управления My Sql
## Дальше очень важный пункт
6. Слева необходимо будет нажать на кнопку **+New** и написать **golang**
Если будете использовать другое приложение для запуска MySql,
нужно будет поменять данный в соотвествии с ваши сервером

6) Запускаем **cmd/app/main.go** (запускаем само приложение)
```rb
go run cmd/app/main.go
```
8) Заходим в консоль и пишем следующую команду
```rb
curl http://127.0.0.1:8080/calculate/?example=8-3*5
```
example указывается пример **также вместо знака + следует вводить %2B**

### Пример:
```rb
curl http://127.0.0.1:8080/calculate/?example=8-3%2B5
```

Будет возвращаться решение примера, а также id самого примера
по id можно смотреть, как решался пример введя след. команду(мониторинг)
```rb
curl http://127.0.0.1:8080/steps/?id=3
```

# Дополнительные примеры
```rb
curl http://127.0.0.1:8080/calculate/?example=5-3*6/3-2
```
```rb
curl http://127.0.0.1:8080/calculate/?example=3*3%2B5-2
```
# Схема проекта 
*Здесь указаны, как взаимодействуют между собой функции, а также описание БД*
![Alt text](https://github.com/tantoni228/server_calculator/blob/main/app_server_calculator.png).

# Недароботки проекта
+ Не может решать примеры со скобками
+ Решение с целыми числами
+ Деление целочислено

# Прошу прощения за ошибки писал в спешке

Вот мой [телеграм](https://t.me/sadlarfox) в случаи вопросов

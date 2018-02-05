# swTestProject

Используются сторонние пакеты 
github.com/gorilla/mux
gopkg.in/mgo.v2
gopkg.in/mgo.v2/bson


В качестве хранилища использовал MongoDb



Тестирование

GET http://localhost:8081 - Ping

GET http://localhost:8081/teams - Список всех команд

GET http://localhost:8081/teams/5a72416a3482f54aa3e1d7e3 - Получение одной команды

POST http://localhost:8081/teams - Создание, пример тела запроса: {"Name": "Бавария"}

PUT http://localhost:8081/team/5a77645635c14ea39ea6a0e4 - Редактирование, пример тела запроса: {"Name": "Севилья"}

DELETE http://localhost:8081/team/5a77645635c14ea39ea6a0e4 - Удаление команды

GET http://localhost:8081/draw/3 - Жеребьёвка, параметром указывается минимальное количество команд в группе

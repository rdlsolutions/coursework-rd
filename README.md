# Coursework. Robot-Dreams

## coursework-rd. Introdution

Для курсового проекту був створений застосунок у вигляді веб-календаря, який зберігає події в БД Postgres (бути Stateful).

![alt text](image/RD-App.png)

В GitHub був створений публічний репозиторій [https://github.com/rdlsolutions/coursework-rd](https://github.com/rdlsolutions/coursework-rd)

### Етап 1. Containerization.

![alt text](image/DockerBuild.png)

Застосунок був збілджений за допомогою Dockerfil з multistage та завантажений до Docker Hub репозиторію [https://hub.docker.com/repository/docker/rdlsolutions/coursework/general](https://hub.docker.com/repository/docker/rdlsolutions/coursework/general)

![alt text](image/DockerHub.png)

ТАкож одразу до завдання був доданий image Postgres:16 для тестування застосунку та побудови Docker-compose.yaml

![alt text](<image/Docker Compose.png>)

### Етап 2. Helm-chart


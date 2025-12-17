# Coursework. Robot-Dreams

## coursework-rd. Introdution

Для курсового проекту був створений застосунок у вигляді веб-календаря, який зберігає події в БД Postgres (стан Stateful).

![alt text](images/RD-App.png)

В GitHub був створений публічний репозиторій [https://github.com/rdlsolutions/coursework-rd](https://github.com/rdlsolutions/coursework-rd)

### Етап 1. Containerization

![alt text](images/DockerBuild.png)

Застосунок був збілджений за допомогою Dockerfil з multistage та завантажений до Docker Hub репозиторію [https://hub.docker.com/repository/docker/rdlsolutions/coursework/general](https://hub.docker.com/repository/docker/rdlsolutions/coursework/general)

![alt text](images/DockerHub.png)

Також одразу до завдання був доданий image Postgres:16 для тестування застосунку та побудови Docker-compose.yaml

![alt text](<images/Docker Compose.png>)

### Етап 2. Helm-chart

#### Flux CD

```sh
  flux bootstrap github --owner=rdlsolutions  --repository=coursework-rd  --branch=main  --path=./clusters/local  --personal
```

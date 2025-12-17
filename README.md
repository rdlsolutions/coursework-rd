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

Небхідно встановити та налаштувати Helm для застосунку.

`helm create coursework`

Налаштовуємо змінні для:

- deployment.yaml
- service.yaml
- ingress.yaml

```sh
  flux bootstrap github --owner=rdlsolutions  --repository=coursework-rd  --branch=main  --path=./clusters/local  --personal
```

### Етап 3. Додаємо БД та оператори

Для застосунку необхідна для роботи БД Postgres, створюємо її через flux kustomize, додаючи до кластеру файл postgres.yaml kind: Cluster від postgresql.cnpg.kustomize.config.k8s.io та додаємо оператор БД CloudNativePG. Паралельно додаємо kustomize ресурси для застосунку.
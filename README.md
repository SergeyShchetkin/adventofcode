Решение задач adventofcode.
---

### Необходимые инструменты для запуска:

1) [Docker](https://www.docker.com/)
2) [docker-compose](https://docs.docker.com/compose/)
3) [make](https://www.gnu.org/software/make/)

### Запуск требуемой задач:

1) Собрать контейнер

```bash
$ make build
```

2) Запустить требуемую задачу

```bash
$ make run-task DIR=day1 TASK=task2
```

* DIR  - каталог с задачами
* TASK - имя задачи

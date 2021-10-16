# Ozon Marketplace Project

![schema](schema.png)

### Задание 1

- Сделать форк **ozonmp/omp-bot** репозитория в свой профиль
- Запросить у своего тьютора свой поддомен **domain/subdomain**
- Добавить поддержку следующих команд

```
/new{subdomain} — create a new entity
/list{subdomain} — get a list of your entity
/edit{subdomain} — edit a entity
/delete{subdomain} — delete an existing entity
```



Для добавления поддержки команд для своей поддомена:

1. Написать структуру `{Subdomain}` с методом `String()`
2. Написать интерфейс `{Subdomain}Service` и **dummy** имплементацию
3. Написать интерфейс `{Subdomain}Commander` по обработке команд

---

2. Написать интерфейс `{Subdomain}Service` в **internal/services/{domain}/{subdomain}.go**

```go
package {domain}

import "github.com/ozonmp/omp-bot/internal/models/domain"

type {Subdomain}Service interface {
  Create({domain}.{Subdomain}) //...🤔
  Update({subdomain}_id uint64, {subdomain} {domain}.{Subdomain}) // ...🤔
  Remove({subdomain}_id uint64) // ...🤔
  List(cursor uint64, limit uint64) // ...🤔
}

type Dummy{Subdomain}Service struct {}

func New() Dummy{Subdomain}Service {
}
```

---

3. Написать интерфейс `{Subdomain}Commander` по обработке команд в **internal/commands/{domain}/{subdomain}.go**

```go
package {domain}

import (
  model "github.com/ozonmp/omp-bot/internal/models/domain"
  service "github.com/ozonmp/omp-bot/internal/services/domain"
)

type {Subdomain}Commander interface {
  New(...)
  List(...)
  Edit(...)
  Delete(...)
}
```


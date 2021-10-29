# Ozon Marketplace Project

![schema](images/schema.png)

Дальше везде используются **placeholder**-ы:

- `{domain}`,`{Domain}`
- `{subdomain}`,`{Subdomain}`

Например, для поддомена `package` из домена `logistic` значение **placeholder**-ов будет:

- `{domain}`,`{Domain}` = `logistic`,`Logistic`
- `{subdomain}`,`{Subdomain}` = `package`,`Package`
- `{domain}`/`{subdomain}` = `logistic`/`package`

---

### Задание 3

1. Сделать **rebase** своего репозитория `{kw-domain}-{subdomain}-api` на [omp-template-api](https://github.com/ozonmp/omp-template-api)
3. Добавить в **proto** следующие **handler**-ы (пример [template](https://github.com/ozonmp/omp-template-api/blob/be1223fb1d1c9751b0d9db1d6e2dfff6ba4c9316/protos/ozonmp/omp_template_api/v1/omp_template_api.proto)):
   1. `Create{Subdomain}`
   2. `Describe{Subdomain}`
   3. `List{Subdomain}s`
   4. `Remove{Subdomain}`
4. Добавить теги валидации в поля сообщений (пример [template](https://github.com/ozonmp/omp-template-api/blob/be1223fb1d1c9751b0d9db1d6e2dfff6ba4c9316/protos/ozonmp/omp_template_api/v1/omp_template_api.proto#L28))
5. Сделать рефакторинг: заменить `template` на `{subomain}` (см. рецепт)
6. Сгенерировать **gRPC** код клиента и сервера (make generate)
7. Имплементировать код новых ручек в **internal/api/api.go** (пример [template](https://github.com/ozonmp/omp-template-api/blob/be1223fb1d1c9751b0d9db1d6e2dfff6ba4c9316/internal/api/api.go#L34))
   1. Код ручек должен просто логгировать вызовы (с уровнем `debug`)
   2. Возвращать пустой ответ или внутреннюю ошибку (`not implemented`)
8. Протестировать через **grpc_cli** написанные ручки (пример [template](https://github.com/ozonmp/omp-template-api/blob/main/DOCS.md#grpc))
9. Написать тесты по обработке не валидных запросов :gem:
10. Сгенерировать **Python** код клиента и задеплоить его в **PyPi** :gem: (пример [template](https://github.com/ozonmp/omp-template-api/blob/main/DOCS.md#python-client))


**Рецепт**

```sh
export domain_kw=omp
export subdomain=demo

git remote add template https://github.com/ozonmp/omp-template-api
git fetch template main
git rebase template/main
git checkout template/main -- Makefile go.mod go.sum
git rebase --continue
rm -rf pkg/omp-template-api
mkdir pkg/${domain_kw}-${subdomain}-api
mv protos/ozonmp/omp_template_api/v1/omp_template_api.proto \
   protos/ozonmp/omp_template_api/v1/${domain_kw}_${subdomain}_api.proto
mv protos/ozonmp/omp_template_api protos/ozonmp/${domain_kw}_${subdomain}_api
mv pypkg/omp-template-api pypkg/${domain_kw}-${subdomain}-api
// grep (exclude 'protos/google' dir)
// - template -> ${subdomain}
// - grep omp -> ${domain_kw}
make generate
go mod tidy
make build
# перенесли в шаблонном репозитории README.md в DOCS.md, чтобы было меньше коонфликтов при rebase
mv DOCS.md README.md
git add .
git commit -m"refactored"
```

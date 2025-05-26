# go-notes-api

Basit bir RESTful Not Tutma API uygulamasıdır.  
Go dili kullanılarak geliştirilmiştir ve Gorilla Mux router ile HTTP isteklerini yönetir.

## Özellikler

- Notları listeleme, oluşturma, silme ve ID ile getirme  
- JSON formatında veri alışverişi  
- UUID tabanlı benzersiz not ID’leri  
- SQLite veritabanı (opsiyonel olarak)

## Kurulum

1.Bu repoyu klonlayın:
 ```bash
   git clone https://github.com/melihkl/go-notes-api.git
   cd go-notes-api
```

## Kullanım

- Sunucu `localhost:8085` portunda çalışır.  
- API endpointleri:  
- `GET /notes` — Tüm notları getirir  
- `POST /notes` — Yeni not oluşturur  
- `GET /notes/{id}` — ID ile not getirir  
- `DELETE /notes/{id}` — ID ile not siler

## Test

Testleri çalıştırmak için:
 ```bash
go test .\tests\
```

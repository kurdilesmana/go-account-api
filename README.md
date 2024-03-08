# API Go

## Docker
Clone this repository and run:
```
docker-compose up
```

Following endpoints:

| Method | Route                | Body                                               |
| ------ | -------------------- |----------------------------------------------------|
| POST   | /daftar              | `{"name": "John Doe","nik": "876372333999","phone_number": "081234567899","pin": "1234"}`|
| POST   | /tabung              | `{"account_number": "00009", "amount": 50000`      |
| POST   | /tarik               | `{"account_number": "00009", "amount": 50000`      |
| POST   | /transfer            | `{"src_account_number": "0812345678990001","dest_account_number": "0812345678900001","amount": 100000}`      |
| GET    | /saldo/:account_number  |                                                 |
| GET    | /mutasi/:account_number |                                                 |

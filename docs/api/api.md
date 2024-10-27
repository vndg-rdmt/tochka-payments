

### Payment

`/api/payment`


```json
{
    "amount": uint64,
    "string": "string"
}
```

```json
{
    "id": "string"
}
```

### Status

`/api/status/{:id}`

```json
{
  "Data": {
    "Operation": [
      {
        "customerCode": "300000092",
        "taxSystemCode": "osn",
        "paymentType": "card",
        "paymentId": "A22031016256670100000533E625FCB3",
        "transactionId": "48232c9a-ce82-1593-3cb6-5c85a1ffef8f",
        "createdAt": "2022-10-18T08:28:59+00:00",
        "paymentMode": [
          "sbp",
          "card"
        ],
        "redirectUrl": "https://example.com",
        "failRedirectUrl": "https://example.com/fail",
        "Client": {
          "name": "Иванов Иван Иванович",
          "email": "ivanov@mail.com",
          "phone": "+7999999999"
        },
        "Items": [
          {
            "vatType": "none",
            "name": "string",
            "amount": "1234.00",
            "quantity": 1,
            "paymentMethod": "full_payment",
            "paymentObject": "service",
            "measure": "шт."
          }
        ],
        "purpose": "Перевод за оказанные услуги",
        "amount": "1234.00",
        "status": "CREATED",
        "operationId": "48232c9a-ce82-1593-3cb6-5c85a1ffef8f",
        "paymentLink": "https://merch.bank24.int/order/?uuid=16ea4c54-bf1d-4e6a-a1ef-53ad55666e43",
        "merchantId": "200000000001056",
        "consumerId": "fedac807-078d-45ac-a43b-5c01c57edbf8"
      }
    ]
  },
  "Links": {
    "self": "https://enter.tochka.com/uapi"
  },
  "Meta": {
    "totalPages": 1
  }
}
```
